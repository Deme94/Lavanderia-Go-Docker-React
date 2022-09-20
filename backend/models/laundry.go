package models

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/mxschmitt/playwright-go"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

// Type Laundry (web scraper)
type Laundry struct {
	url               string
	user              string
	password          string
	stripeKey         string
	GoogleLoginClient string
}

func newLaundry(laundry_url string, user string, password string, stripeKey string, googleLoginClient string) *Laundry {

	err := playwright.Install()
	if err != nil {
		log.Fatalf("Error installing playwright: %v", err)
	}

	l := Laundry{
		url:               laundry_url,
		user:              user,
		password:          password,
		stripeKey:         stripeKey,
		GoogleLoginClient: googleLoginClient,
	}
	return &l
}

func (l *Laundry) GetWashers() (map[string]*Machine, error) {
	washers, err := l.getMachines("LAVADORA")
	if err != nil {
		return nil, err
	}

	return washers, nil
}

func (l *Laundry) GetDryers() (map[string]*Machine, error) {
	dryers, err := l.getMachines("SECADORA")
	if err != nil {
		return nil, err
	}

	return dryers, nil
}

func (l *Laundry) CreatePaymentIntent(amount int64) (string, error) {

	stripe.Key = l.stripeKey

	// Create a PaymentIntent with amount and currency
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),                      // va en centimos
		Currency: stripe.String(string(stripe.CurrencyEUR)), // va en centimos
		PaymentMethodTypes: stripe.StringSlice([]string{
			*stripe.String(string(stripe.PaymentMethodTypeCard)),
		}),
		ReceiptEmail: stripe.String("deme1994@gmail.com"),
	}

	pi, err := paymentintent.New(params)

	if err != nil {
		return "", err
	}

	return pi.ClientSecret, nil
}

func (l *Laundry) ActivateMachine(IDMachine int) error {
	return nil
}

func (l *Laundry) getMachines(machineType string) (map[string]*Machine, error) {
	// Launch browser
	pw, err := playwright.Run()
	if err != nil {
		errString := fmt.Sprintf("could not start playwright: %v", err)
		err = errors.New(errString)
		return nil, err
	}
	var headless = true
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Args: []string{"--disable-gpu"}, Headless: &headless})
	if err != nil {
		errString := fmt.Sprintf("could not launch browser: %v", err)
		err = errors.New(errString)
		return nil, err
	}
	defer browser.Close()
	defer pw.Stop()
	page, err := browser.NewPage()
	if err != nil {
		errString := fmt.Sprintf("could not create page: %v", err)
		err = errors.New(errString)
		return nil, err
	}
	// Go to web
	_, err = page.Goto(l.url, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	})
	if err != nil {
		errString := fmt.Sprintf("could not goto: %v", err)
		err = errors.New(errString)
		return nil, err
	}
	// LOGIN PAGE
	// Obtain login frame
	el, err := page.WaitForSelector("html > frameset > frame:nth-child(2)")
	if err != nil {
		errString := fmt.Sprintf("frame not found: %v", err)
		err = errors.New(errString)
		return nil, err
	}
	frame, err := el.ContentFrame()
	if err != nil {
		errString := fmt.Sprintf("element is not a frame: %v", err)
		err = errors.New(errString)
		return nil, err
	}
	// Login: user > psswd > accept
	err = frame.Click("#txtusuario")
	if err != nil {
		errString := fmt.Sprintf("could not click: %v", err)
		err = errors.New(errString)
		return nil, err
	}
	err = frame.Type("#txtusuario", l.user)
	if err != nil {
		errString := fmt.Sprintf("could not type: %v", err)
		err = errors.New(errString)
		return nil, err
	}
	err = frame.Click("#txtContraseña")
	if err != nil {
		errString := fmt.Sprintf("could not click: %v", err)
		err = errors.New(errString)
		return nil, err
	}
	err = frame.Type("#txtContraseña", l.password)
	if err != nil {
		errString := fmt.Sprintf("could not type: %v", err)
		err = errors.New(errString)
		return nil, err
	}
	err = frame.Press("#btnAcpetar", "Enter")
	if err != nil {
		errString := fmt.Sprintf("could not create press: %v", err)
		err = errors.New(errString)
		return nil, err
	}

	// HOME PAGE
	frame.Click("#lblMenu > ul > li:nth-child(2) > a")

	// CONFIGURATION
	//frame.WaitForTimeout(2000)
	frame.WaitForLoadState()
	values := []string{"25"}
	frame.SelectOption("#example_length > label > span.custom-select > select", playwright.SelectOptionValues{Values: &values})
	//frame.WaitForTimeout(2000)

	names, err := frame.QuerySelectorAll("tbody > tr > td:nth-child(1)")
	if len(names) == 0 {
		errString := "could not get entries"
		err = errors.New(errString)
		return nil, err
	}
	if err != nil {
		errString := fmt.Sprintf("could not get entries: %v", err)
		err = errors.New(errString)
		return nil, err
	}
	ids, err := frame.QuerySelectorAll("tbody > tr > td:nth-child(3)")
	if err != nil {
		errString := fmt.Sprintf("could not get entries: %v", err)
		err = errors.New(errString)
		return nil, err
	}
	prices, err := frame.QuerySelectorAll("tbody > tr > td:nth-child(4)")
	if err != nil {
		errString := fmt.Sprintf("could not get entries: %v", err)
		err = errors.New(errString)
		return nil, err
	}

	machinesMap := make(map[string]*Machine)
	for i := 0; i < len(names); i++ {
		name, err := names[i].TextContent()
		if err != nil {
			errString := fmt.Sprintf("could not get text content: %v", err)
			err = errors.New(errString)
			return nil, err
		}
		if strings.Contains(name, machineType) {
			idString, err := ids[i].TextContent()
			if err != nil {
				errString := fmt.Sprintf("could not get text content: %v", err)
				err = errors.New(errString)
				return nil, err
			}
			priceString, err := prices[i].TextContent()
			if err != nil {
				errString := fmt.Sprintf("could not get text content: %v", err)
				err = errors.New(errString)
				return nil, err
			}

			id, err := strconv.Atoi(idString)
			if err != nil {
				return nil, err
			}
			priceString = strings.Split(priceString, " ")[0]
			priceString = strings.Replace(priceString, ",", ".", 1)
			price, err := strconv.ParseFloat(priceString, 64)
			if err != nil {
				return nil, err
			}

			m := Machine{
				ID:       id,
				Name:     name,
				Status:   "",
				TimeLeft: 0,
				Price:    price,
			}
			machinesMap[name] = &m
		}
	}

	// Go to machines state
	frame.Click("#lblMenu > ul > li:nth-child(5) > a")

	// MACHINES STATE
	frame.WaitForLoadState()
	frame.WaitForSelector("#form1")
	//frame.WaitForTimeout(2000)

	names, err = frame.QuerySelectorAll("div > div.pmd-display2")
	if len(names) == 0 {
		errString := "could not get entries"
		err = errors.New(errString)
		return nil, err
	}
	if err != nil {
		errString := fmt.Sprintf("could not get entries: %v", err)
		err = errors.New(errString)
		return nil, err
	}
	startTimes, err := frame.QuerySelectorAll("div > div.source-semibold.typo-fill-secondary")
	if err != nil {
		errString := fmt.Sprintf("could not get entries: %v", err)
		err = errors.New(errString)
		return nil, err
	}
	statusImgs, err := frame.QuerySelectorAll("div > a > img")
	if err != nil {
		errString := fmt.Sprintf("could not get entries: %v", err)
		err = errors.New(errString)
		return nil, err
	}

	for i := 0; i < len(names); i++ {
		name, err := names[i].TextContent()
		if err != nil {
			errString := fmt.Sprintf("could not get text content: %v", err)
			err = errors.New(errString)
			return nil, err
		}
		if strings.Contains(name, machineType) {

			startTime, err := startTimes[i].TextContent()

			if err != nil {
				errString := fmt.Sprintf("could not get text content: %v", err)
				err = errors.New(errString)
				return nil, err
			}
			var timeLeft int
			if startTime != "-" {
				startTimeSplitted := strings.Split(startTime, ":")
				startTimeHours, err := strconv.Atoi(startTimeSplitted[0])
				if err != nil {
					return nil, err
				}
				startTimeMinutes, err := strconv.Atoi(startTimeSplitted[1])
				if err != nil {
					return nil, err
				}
				startTimeDec := startTimeHours + startTimeMinutes/60.0
				currentTime := time.Now().Format("15:04")
				currentTimeSplitted := strings.Split(currentTime, ":")
				currentTimeHours, err := strconv.ParseFloat(currentTimeSplitted[0], 64)
				if err != nil {
					return nil, err
				}
				currentTimeMinutes, err := strconv.ParseFloat(currentTimeSplitted[1], 64)
				if err != nil {
					return nil, err
				}
				currentTimeDec := currentTimeHours + currentTimeMinutes/60.0
				timeLeft = int((30 - (currentTimeDec - float64(startTimeDec))) * 60)
			} else {
				timeLeft = 0
			}

			statusImg, err := statusImgs[i+1].GetAttribute("src")
			if err != nil {
				errString := fmt.Sprintf("could not get attribute: %v", err)
				err = errors.New(errString)
				return nil, err
			}

			var status string
			switch statusImg {
			case "/imgs/lavlibre.png":
				status = "green"
			case "/imgs/lavocup.png":
				status = "red"
			case "/imgs/lavpend.png":
				status = "blue"
			case "/imgs/seclibre.png":
				status = "green"
			case "/imgs/secocup.png":
				status = "red"
			case "/imgs/secpend.png":
				status = "blue"
			default:
				log.Println("NO CARGA LA IMAGEN DE LA MAQUINA")
			}

			machinesMap[name].TimeLeft = timeLeft
			machinesMap[name].Status = status
		}
	}

	return machinesMap, nil
}
