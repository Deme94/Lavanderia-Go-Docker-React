package models

import (
	"os"
)

// CONSTANTS -----------------------------------------

// Database Postgres SQL
var dbDriverName = os.Getenv("DB_DRIVER_NAME")
var dbIP = os.Getenv("DB_IP")
var dbPort = os.Getenv("DB_PORT")
var dbSslMode = os.Getenv("DB_SSL_MODE")
var dbUser = os.Getenv("POSTGRES_USER")
var dbPswd = os.Getenv("POSTGRES_PASSWORD")
var dbName = os.Getenv("POSTGRES_DB")

// Laundry
var laundryUrl = os.Getenv("LAUNDRY_URL")
var laundryUser = os.Getenv("LAUNDRY_USER")
var laundryPswd = os.Getenv("LAUNDRY_PASSWORD")
var stripeKey = os.Getenv("STRIPE_KEY")
var googleLoginClient = os.Getenv("GOOGLE_LOGIN_CLIENT")

// ---------------------------------------------------

// Models is the wrapper for models
type Models struct {
	DBPostgres *DatabaseSQL
	Laundry    *Laundry
}

// NewModels returns all models
func NewModels() *Models {
	return &Models{
		DBPostgres: newDatabaseSQL(dbDriverName, dbDriverName+"://"+dbUser+":"+dbPswd+"@"+dbIP+":"+dbPort+"/"+dbName+"?sslmode="+dbSslMode),
		Laundry:    newLaundry(laundryUrl, laundryUser, laundryPswd, stripeKey, googleLoginClient),
	}
}
