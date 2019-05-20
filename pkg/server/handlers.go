package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/database"
	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/hoovermap"
)

// IndexPost is the method that will be called when the route '/' is POSTed too.
func IndexPost(w http.ResponseWriter, r *http.Request) {
	// Create a Timestamp for this request that we will use to persist all data
	// within this operation.
	// Values look like:
	//             2019-05-20 14:09:30.709840974 +0100 BST m=+3.073307183
	//
	// Chances of collision are very low, string is a bit ugly, it could be hashed
	// but that seems like it would obscure anyone attempting to understand its purpose.
	// That purpose is to identify which Inputs relate to which Outputs
	// and keeping the key name as the time allows this matching to be performed and
	// time stamped.
	//
	// Alternatively, the time could be stored along with the I/O objects with additional
	// values - but this was not in the spec.
	key := []byte(time.Now().String())

	// Read the request body.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("500 could not deserialise the request body: " + err.Error()))
		return
	}

	// Attempt to deserialise the request into an Input struct.
	var input hoovermap.APIInput
	err = json.Unmarshal(body, &input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 could not understand the request body: " + err.Error()))
		return
	}

	// Write the input, we have deserialised it already, so it should at be valid JSON.
	err = database.WriteInput(key, body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 could not write database value: " + err.Error()))
		return
	}

	// Compute the movement and return a response struct.
	output, err := hoovermap.ComputeMovement(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	// Serialise the output to byte encoded JSON.
	outputBytes, err := json.Marshal(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 could not serialise the response:" + err.Error()))
		return
	}

	// Write the output.
	err = database.WriteOutput(key, outputBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 could not write database value: " + err.Error()))
		return
	}

	// Return the computed response.
	w.WriteHeader(http.StatusOK)
	w.Write(outputBytes)
}

// Index is a simple HelloWorld endpoint.
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
