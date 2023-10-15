package forms

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// Form  creates custom form struct and it embeds url.values object
type Form struct {
	url.Values
	Errors errors
}

// New initializes form structure
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Valid() bool {
	log.Println("this is", f.Errors)

	return len(f.Errors) == 0
}

// Has Checks  form filed is not empty.
func (f *Form) Has(field string, r *http.Request) bool {
	log.Println(field)
	x := r.Form.Get(field)
	log.Println(x)
	if x == "" {
		name := "First Name"
		f.Errors.Add(field, fmt.Sprintf("%s can't be empty", name))
		return false
	}
	return true
}
