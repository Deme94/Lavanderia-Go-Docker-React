package main

import (
	"Lavanderia/models"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/natefinch/lumberjack"
)

// CONSTANTS ----------------------------------------

// App version
const VERSION = "1.0.0"

var domain = os.Getenv("DOMAIN")

// --------------------------------------------------

type config struct {
	port int
	env  string
	jwt  struct {
		secret string
	}
}

type application struct {
	config config
	logger *log.Logger
	models *models.Models
}

func main() {
	// Config
	p, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
	env := os.Getenv("ENVIRONMENT")

	var cfg config
	flag.IntVar(&cfg.port, "port", p, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", env, "Application environment ("+env+")")
	flag.Parse()

	cfg.jwt.secret = os.Getenv("SERVER_JWT")

	// Logger
	os.Mkdir("./log", os.ModePerm)
	filename := "./log/logs.txt"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, "", log.Ldate|log.Ltime|log.LstdFlags|log.Lshortfile)

	logger.SetOutput(&lumberjack.Logger{
		Filename: filename,
		MaxSize:  1, // megabytes after which new file is created
		//MaxBackups: 3,  // number of backups
		//MaxAge:     28, //days
	})

	// Models
	models := models.NewModels()

	// Server application
	app := &application{
		config: cfg,
		logger: logger,
		models: models,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	srv.SetKeepAlivesEnabled(true)

	logger.Println("Started server on port", cfg.port)
	log.Println("Started server on port", cfg.port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
