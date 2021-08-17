package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nstoker/fictional-pancake/internal/handlers"
	"github.com/nstoker/fictional-pancake/internal/version"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("Hello, world! %s starting", version.Version())

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Infof("error loading godotenv file %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3030"
		logrus.Infof("environment variable PORT not set, defaulting to %s", port)
	}

	frontend := os.Getenv("FRONTEND")
	if frontend == "" {
		frontend = "./frontend/dist"
		logrus.Infof("environment variable FRONTEND not set, defaulting to %s", frontend)
	}

	logrus.Infof("serving spa from %s", frontend)

	http.HandleFunc("/hello", handlers.HomePageHandler)

	fs := http.FileServer(http.Dir(frontend))
	http.Handle("/", fs)

	logrus.Infof("Listening on %s", port)

	logrus.Panic(
		http.ListenAndServe(fmt.Sprintf(":%s", port), nil),
	)
}
