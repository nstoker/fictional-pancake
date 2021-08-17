package handlers

import (
	"fmt"
	"net/http"

	"github.com/nstoker/fictional-pancake/internal/version"
	"github.com/sirupsen/logrus"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, World! %s", version.Version())
	if err != nil {
		logrus.Panic(err)
	}
}
