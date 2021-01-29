package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/DamienBirtel/FizzBuzz/lib"
)

// FizzBuzzHandler is a http.Handler
type FizzBuzzHandler struct {
}

// NewFizzBuzzHandler returns a FizzBuzz Handler
func NewFizzBuzzHandler() *FizzBuzzHandler {
	return &FizzBuzzHandler{}
}

// ServeHTTP satisfies the http.Handler interface. It creates a lib.FizzBuzz struct and calls
// it's function Print.
func (f *FizzBuzzHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	// we only handle GET and POST requests
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		fmt.Fprintf(rw, "Only GET and POST methods are supported at the moment")
		return
	}

	// we create a struct with parameters by default
	fb := &lib.FizzBuzz{
		Length:   200,
		FizzNum:  7,
		BuzzNum:  9,
		FizzWord: "Fizz",
		BuzzWord: "Buzz",
	}

	// if there is a request body, it's assumed to be containing a JSON with custom parameters
	// which we decode into the FizzBuzz struct
	if r.Method == http.MethodPost {
		err := fb.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Error reading FizzBuzz custom parameters", http.StatusBadRequest)
			return
		}
	}

	fb.Print(os.Stdout)
}
