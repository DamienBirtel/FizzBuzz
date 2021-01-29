package lib

import (
	"encoding/json"
	"fmt"
	"io"
)

// FizzBuzz contains all the information needed to print a custom fizzbuzz
type FizzBuzz struct {
	Length   int    `json:"length,omitempty"`
	FizzNum  int    `json:"fizznum,omitempty"`
	BuzzNum  int    `json:"buzznum,omitempty"`
	FizzWord string `json:"fizzword,omitempty"`
	BuzzWord string `json:"buzzword,omitempty"`
}

func isMultiple(testedNumber int, rootNumber int) bool {
	if testedNumber%rootNumber == 0 {
		return true
	}
	return false
}

// FromJSON decodes json from an io.Reader to FizzBuzz{}
func (f *FizzBuzz) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(f)
}

// Print writes numbers from 1 to f.Length to an io.Writer,
// replacing every multiple of f.FizzNum by f.FizzWord, every multiple of f.BuzzNum by f.BuzzWord
// and every multiple of both by the concatenation of f.FizzWord and f.BuzzWord
func (f *FizzBuzz) Print(w io.Writer) {

	fizzBuzzNum := f.FizzNum * f.BuzzNum
	fizzBuzzWord := f.FizzWord + f.BuzzWord

	for i := 1; i <= f.Length; i++ {
		if isMultiple(i, fizzBuzzNum) {
			fmt.Fprintf(w, "%s\n", fizzBuzzWord)
		} else if isMultiple(i, f.FizzNum) {
			fmt.Fprintf(w, "%s\n", f.FizzWord)
		} else if isMultiple(i, f.BuzzNum) {
			fmt.Fprintf(w, "%s\n", f.BuzzWord)
		} else {
			fmt.Fprintf(w, "%d\n", i)
		}
	}
}
