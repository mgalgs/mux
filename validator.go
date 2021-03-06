package mux

import "net/http"

//Validator validates the incomming value against a valid value/s
type Validator interface {
	Validate(RouteInterface) error
}

//MethodValidator validates the string against a method.
type MethodValidator map[string]struct{}

// newMethodValidator returns default method validator
func newMethodValidator() MethodValidator {
	return MethodValidator(methods)
}

// methods all possible standard methods
var methods = map[string]struct{}{
	http.MethodGet:     struct{}{},
	http.MethodPost:    struct{}{},
	http.MethodPatch:   struct{}{},
	http.MethodDelete:  struct{}{},
	http.MethodHead:    struct{}{},
	http.MethodPut:     struct{}{},
	http.MethodOptions: struct{}{},
	http.MethodConnect: struct{}{},
}

func (v MethodValidator) Validate(r RouteInterface) error {

	if _, found := v[r.GetMethodName()]; !found {
		return NewBadMethodError(r.GetMethodName())
	}

	return nil
}

//pathMatcherValidator validates the string against a method.
type pathMatcherValidator struct{}

func newPathMatcherValidator() pathMatcherValidator {
	return pathMatcherValidator{}
}

func (v pathMatcherValidator) Validate(r RouteInterface) error {

	for _, m := range r.GetMatchers() {
		if m.Rank() == rankPath {
			return nil
		}
	}

	return NewMissingPathError()
}
