package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HandleTest(t *testing.T) {
	t.Run("should return an error if authentication fails", func(t *testing.T) {
		app := app{
			Username: "Username",
		}

		request, _ := http.NewRequest("POST", "", nil)
		responseRecorder := httptest.NewRecorder()

		app.handle(responseRecorder, request)

		assert.Equal(t, 401, responseRecorder.Code)
	})

	t.Run("should return an error if unmarshalling fails", func(t *testing.T) {
		// we need make the authorization step pass -> configure the app, put valid user into the request
		// provide a request object that cannot be parsed
	})

	t.Run("should return an error if validation fails", func(t *testing.T) {
		// we need to make the authorization step pass -> configure the app, put valid user into the request
		// we need to make the parsing step pass
		// provide an invalid request object
	})

	t.Run("should return an error if transformation", func(t *testing.T) {
		// we need to make the authorization step pass -> configure the app, put valid user into the request
		// we need to make the parsing step pass
		// we need to make the validation step pass
		// provide invalid configuration
	})

	// ... and so on
}

func TestAuthorize(t *testing.T) {}

func TestUnmarshall(t *testing.T) {}

func TestValidate(t *testing.T) {}

func TestTransform(t *testing.T) {}

func TestCallExtApi(t *testing.T) {}

func TestRespondWithOK(t *testing.T) {}

func TestRspondWithInternalServicerError(t *testing.T) {}
