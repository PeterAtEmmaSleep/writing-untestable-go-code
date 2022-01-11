package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type app struct {
	Username  string
	SomeField string
}

func (app *app) handle(w http.ResponseWriter, req *http.Request) {
	_ = app.authorize(req)
	appRequest, _ := app.unmarshall(req)
	_ = app.validate(appRequest)
	extRequest, _ := app.transform(appRequest)
	_ = app.callExtApi(extRequest)
	_ = app.respondWithOK(w)
}

func (app *app) authorize(req *http.Request) error {
	username := req.Header.Get("username")
	if username == app.Username {
		return nil
	}
	return errors.New("user is not authorized to perform request")
}

func (app *app) unmarshall(httpRequest *http.Request) (AppRequest, error) {
	defer httpRequest.Body.Close()
	var request AppRequest
	body, _ := io.ReadAll(httpRequest.Body)
	_ = json.Unmarshal(body, &request)
	return request, nil
}

func (app *app) validate(appRequest AppRequest) error {
	return appRequest.validate()
}

func (app *app) transform(appRequest AppRequest) (ExtRequest, error) {
	return ExtRequest{someField: appRequest.someField + app.SomeField}, nil
}

func (app *app) callExtApi(_ ExtRequest) error {
	return nil
}

func (app *app) respondWithOK(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusOK)
	return nil
}

func (app *app) respondWithInternalServicerError(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusInternalServerError)
	return nil
}
