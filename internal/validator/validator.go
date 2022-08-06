package validator

import "regexp"

var EmailRX = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

// Validator new validator with a map of validation errors
type Validator struct {
	Errors map[string]string
}

// New is the helper that creates the validator instance
// with an empty errors map
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

// Valid checks if the validator is valid
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// AddError adds an error to the errors map
func (v *Validator) AddError(key, msg string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = msg
	}
}

// Check an error msg only if ok is true
func (v *Validator) Check(ok bool, key, msg string) {
	if !ok {
		v.AddError(key, msg)
	}
}

func In(value string, list ...string) bool {
	for i := range list {
		if list[i] == value {
			return true
		}
	}
	return false
}

func Matches(needle string, rx *regexp.Regexp) bool {
	return rx.MatchString(needle)
}

func Unique(values []string) bool {
	uniqueValues := make(map[string]bool)
	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(uniqueValues) == len(values)
}
