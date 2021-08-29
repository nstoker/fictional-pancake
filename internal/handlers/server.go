package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	log "github.com/nstoker/fictional-pancake/internal/app_logger"
)

type thumbnailRequest struct {
	Url string `json:"url"`
}

type screenshotAPIRequest struct {
	Token          string `json:"token"`
	Url            string `json:"url"`
	Output         string `json:"output"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	ThumbnailWidth int    `json:"thumbnail_width"`
	FailOnError    bool   `json:"fail_on_error"`
}

func ThumbnailHandler(w http.ResponseWriter, r *http.Request) {
	var decoded thumbnailRequest

	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	apiKey := os.Getenv("SCREENSHOT_API_KEY")
	if apiKey == "" {
		log.Logger.Errorf("environment variable 'SCREENSHOT_API_KEY' missing")
		http.Error(w, "environment error", http.StatusInternalServerError)
	}
	apiRequest := screenshotAPIRequest{
		Token:          apiKey,
		Url:            decoded.Url,
		Output:         "json",
		Width:          1920,
		Height:         1080,
		ThumbnailWidth: 300,
		FailOnError:    true,
	}
	jsonString, err := json.Marshal(apiRequest)
	if err != nil {
		log.Logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest("POST",
		"https://screenshotapi.net/api/v1/screenshot",
		bytes.NewBuffer(jsonString),
	)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer response.Body.Close()

	if err != nil {
		log.Logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type screenshotAPIResponse struct {
		Screenshot string `json:"screenshot"`
	}
	var apiResponse screenshotAPIResponse
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	if err != nil {
		log.Logger.Error(err)

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = fmt.Fprintf(w, `{ "screenshot": "%s" }`, apiResponse.Screenshot)
	if err != nil {
		log.Logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Logger.Infof("Got the following URL %s", decoded.Url)
}