package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"paste-corral/data"
)

// HandleRequest is the main handler function.
func HandleRequest(w http.ResponseWriter, r *http.Request) {

	var err error

	// -- -----------------------------------
	// At the moment only GET is supported.
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	}

	// -- -----------------------------------
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// -- ---------------------------------------
// Handle GET request
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {

	// -- -----------------------------------
	// Get single post from DB and put into struct:
	pasteResp, err := data.Pastes()

	if err != nil {
		fmt.Println(err)
		return
	}

	// -- -----------------------------------
	// Marshal the struct into JSON:
	output, err := json.MarshalIndent(&pasteResp, "", "\t\t")

	if err != nil {
		fmt.Println(err)
		return
	}

	// -- -----------------------------------
	// Write JSON to the ResponseWriter:
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	return
}
