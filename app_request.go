package main

import "errors"

type AppRequest struct {
	someField string
}

func (request *AppRequest) validate() error {
	if request.someField == "" {
		return errors.New("invalid request")
	}
	return nil
}
