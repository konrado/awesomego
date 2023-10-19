package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Declare a string containing the application version number.
const version = "1.0.0"

// Config struct to hold all the configuration settings for our application
type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Initialize a new logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Declare an instance of the application struct, containing the config struct and the logger
	app := &application{
		config: cfg,
		logger: logger,
	}
	r := app.router()

	_ = r.Run(fmt.Sprintf(":%d", app.config.port))
	fmt.Println("Hello world!")
}
