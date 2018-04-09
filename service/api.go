package service

import (
	"encoding/json"
	"net/http"

	"github.com/aryahadii/runandeh/runner"

	"github.com/aryahadii/runandeh/configuration"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

// StartAPI starts listening for http requests
func StartAPI() {
	address := configuration.GetInstance().GetString("addr")

	router := mux.NewRouter()
	router.HandleFunc("/run", runHandler).Methods("POST")
	logrus.Fatal(http.ListenAndServe(address, router))
}

func runHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Decode request's body
	runRequest := &runner.RunRequest{}
	if err := json.NewDecoder(r.Body).Decode(runRequest); err != nil {
		logrus.WithError(err).Error("can't decode run request's body")
		return
	}

	response, err := runner.Run(runRequest)
	if err != nil {
		logrus.WithError(err).Error("can't run")
	}
	json.NewEncoder(w).Encode(*response)
}
