package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form creates a custom form struct, embeds a url.Value strcut
type Form struct {
	Data   url.Values
	Errors errors
}

// Valid returns true if there are no error, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// check for required field
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Data.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string) bool {
	x := f.Data.Get(field)
	return x != ""
}

// MinLength checks for string minimun length
func (f *Form) MinLength(field string, length int) bool {
	x := f.Data.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("this field must be at least %d characters long", length))
		return false
	}
	return true
}

// IsEmail checks for valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Data.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
