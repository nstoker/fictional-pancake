package main

import (
	"net/http"
	"os"

	"github.com/nstoker/fictional-pancake/internal/handlers"
	"github.com/nstoker/fictional-pancake/internal/version"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("Hello, world! %s starting", version.Version())

	port := os.Getenv("PORT")

	http.HandleFunc("/", handlers.HomePageHandler)

	logrus.Infof("Listening on %s", port)

	logrus.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
