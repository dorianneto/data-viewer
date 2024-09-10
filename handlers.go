package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/vansante/go-ffprobe.v2"
)

type InputPayload struct {
	Link string `json:"link"`
}

type OutputPayload struct {
	Filename string            `json:"filename"`
	Streams  []*ffprobe.Stream `json:"streams"`
}

func (app *application) metadataHandler(w http.ResponseWriter, r *http.Request) {
	var payload InputPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	data, err := ffprobe.ProbeURL(context.TODO(), payload.Link)
	if err != nil {
		app.logger.Error(fmt.Sprintf("Error getting data: %v", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	parsedData := OutputPayload{
		Filename: data.Format.Filename,
		Streams:  data.Streams,
	}

	output, _ := json.Marshal(parsedData)

	w.Header().Add("Content-Type", "application/json")

	_, err = w.Write(output)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("okay!"))
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
