package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/nstoker/fictional-pancake/internal/app_logger"
	"github.com/nstoker/fictional-pancake/internal/handlers"
	"github.com/nstoker/fictional-pancake/internal/version"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	defer log.Logger.Sync()
	log.Logger.Infof("Hello, world! %s starting", version.Version())

	err := godotenv.Load(".env")
	if err != nil {
		log.Logger.Infof("loading godotenv: %v", err)
	}

	port, frontend := loadEnvironmentVariables()

	log.Logger.Infof("serving spa from %s", frontend)

	r := mux.NewRouter()
	r.HandleFunc("/api/health", handlers.HealthHandler).Methods("GET")

	spa := spaHandler{staticPath: frontend, indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)
	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	host := fmt.Sprintf("localhost:%s", port)
	log.Logger.Infof("Listening on %s", host)

	walkRoutes(r)

	r.Use(loggingMiddleware)

	srv := &http.Server{
		Handler:      r,
		Addr:         host,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Logger.Fatalf("%v", srv.ListenAndServe())
}

func loadEnvironmentVariables() (string, string) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3030"
		log.Logger.Infof("environment variable PORT not set, defaulting to %s", port)
	}

	frontend := os.Getenv("FRONTEND")
	if frontend == "" {
		frontend = "./frontend/dist"
		log.Logger.Infof("environment variable FRONTEND not set, defaulting to %s", frontend)
	}

	return port, frontend
}

func walkRoutes(r *mux.Router) {
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Logger.Info(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
