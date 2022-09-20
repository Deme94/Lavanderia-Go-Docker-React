package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/pascaldekloe/jwt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"
)

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

//statusHandler handles /status
func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := AppStatus{
		Status:      "Available",
		Environment: app.config.env,
		Version:     VERSION,
	}
	app.writeJSON(w, http.StatusOK, currentStatus, "status")
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//login handles /login
func (app *application) login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	validUser, err := app.models.DBPostgres.GetUser(creds.Email)
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	hashedPassword := validUser.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	// Generate jwt token after successful login
	jwtToken, err := generateJwtToken(fmt.Sprint(validUser.ID), app.config.jwt.secret)
	if err != nil {
		app.errorJSON(w, err, http.StatusNotImplemented)
		return
	}

	app.writeJSON(w, http.StatusOK, string(jwtToken), "response")
}

type GoogleCredentials struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	GoogleId string `json:"googleId"`
}

//login handles /login
func (app *application) loginGoogle(w http.ResponseWriter, r *http.Request) {
	var creds GoogleCredentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_, err = idtoken.Validate(context.Background(), creds.Token, app.models.Laundry.GoogleLoginClient)
	if err != nil {
		app.errorJSON(w, errors.New("error google login validation"))
		return
	}

	// Generate jwt token after successful login
	jwtToken, err := generateJwtToken(creds.GoogleId, app.config.jwt.secret)
	if err != nil {
		app.errorJSON(w, err, http.StatusNotImplemented)
		return
	}

	app.writeJSON(w, http.StatusOK, string(jwtToken), "response")
}

//register handles /register
func (app *application) register(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 12)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.models.DBPostgres.CreateUser(creds.Email, string(hashedPassword))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusOK, "user created successfully", "response")
}

//getWashingMachines handles /washers
func (app *application) getWashers(w http.ResponseWriter, r *http.Request) {
	washers, err := app.models.Laundry.GetWashers()
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		app.logger.Println(err)
		return
	}

	app.writeJSON(w, http.StatusOK, washers, "washers")
}

//getDryers handles /dryers
func (app *application) getDryers(w http.ResponseWriter, r *http.Request) {
	dryers, err := app.models.Laundry.GetDryers()
	if err != nil {
		app.errorJSON(w, errors.New("could not get machines"), http.StatusInternalServerError)
		app.logger.Println(err)
		return
	}

	app.writeJSON(w, http.StatusOK, dryers, "dryers")
}

type paymentIntentPayload struct {
	Amount int64 `json:"amount"`
}

//createPaymentIntent handles /createPaymentIntent
func (app *application) createPaymentIntent(w http.ResponseWriter, r *http.Request) {

	var pay paymentIntentPayload
	err := json.NewDecoder(r.Body).Decode(&pay)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	pay.Amount = pay.Amount * 100
	if pay.Amount < 50 {
		app.logger.Println("¡Cantidad de pago es menor a 50 céntimos!")
		app.errorJSON(w, errors.New("payment denied"))
		return
	}

	clientSecret, err := app.models.Laundry.CreatePaymentIntent(pay.Amount)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	app.writeJSON(w, http.StatusOK, struct {
		ClientSecret string `json:"clientSecret"`
	}{
		ClientSecret: clientSecret,
	}, "paymentIntent")
}

type activateMachinePayload struct {
	IDMachine int64 `json:"id"`
}

//confirmPayment handles /confirmPayment
func (app *application) confirmPayment(w http.ResponseWriter, r *http.Request) {
	var payload activateMachinePayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorJSON(w, errors.New("error activating machine"), http.StatusInternalServerError)
		app.logger.Fatalf("Error decoding json activateMachinePayload -> %v", err)
		return
	}

	err = app.models.Laundry.ActivateMachine(int(payload.IDMachine))
	if err != nil {
		app.errorJSON(w, errors.New("error activating machine"), http.StatusInternalServerError)
		app.logger.Fatalf("Error activating machine -> %v", err)
		return
	}
	// activar maquina
	app.logger.Printf("PAGO EXITOSO --> maquina %d activada", payload.IDMachine)
	app.writeJSON(w, http.StatusOK, "payment received -> machine activated successfully", "payment")
}

//generateJwtToken generates jwt token
func generateJwtToken(subject string, secret string) ([]byte, error) {

	var claims jwt.Claims
	claims.Subject = fmt.Sprint(subject)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = domain
	claims.Audiences = []string{domain}

	jwtToken, err := claims.HMACSign(jwt.HS256, []byte(secret))
	if err != nil {
		return nil, err
	}

	return jwtToken, nil
}
