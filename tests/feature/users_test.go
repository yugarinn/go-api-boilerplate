package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yugarinn/go-api-boilerplate/tests/mocks"
	"github.com/yugarinn/go-api-boilerplate/app/users/factories"
)


type ExpectedUserResponse struct {
	ID uint64
	Email string
	Name string
	LastName string
	CountryCode string
	PhoneNumber string
	PhonePrefix string
}

func TestUsers(t *testing.T) {
	t.Run("POST /users creates a new user", func(t *testing.T) {
		ResetDatabase()

		var response ExpectedUserResponse
		var payload = []byte(`{"phonePrefix":"+34", "phoneNumber":"123456789", "countryCode":"ES"}`)

		_, router := SetupRouter()
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))

		router.ServeHTTP(writer, request)
		json.NewDecoder(writer.Body).Decode(&response)

		assert.Equal(t, 201, writer.Code)
		assert.Equal(t, "+34", response.PhonePrefix)
		assert.Equal(t, "123456789", response.PhoneNumber)
		assert.Equal(t, "ES", response.CountryCode)
		assert.Equal(t, true, DatabaseHas("users_users", "phone_prefix='+34' AND phone_number='123456789' AND country_code='ES'"))
	})

	t.Run("POST /users creates a new user validation", func(t *testing.T) {
		ResetDatabase()

		_, router := SetupRouter()

		var payload = []byte(`{"phonePrefix":"+34", "phoneNumber":"123456789", "countryCode":"ES"}`)

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))

		router.ServeHTTP(writer, request)

		assert.Equal(t, 201, writer.Code)
		assert.Equal(t, true, DatabaseHas("users_validations", "user_id='1'"))
	})

	t.Run("POST /users sends a verification SMS upon successfully creating a new user", func(t *testing.T) {
		ResetDatabase()

		var response ExpectedUserResponse

		app, router := SetupRouter()

		var payload = []byte(`{"phonePrefix":"+34", "phoneNumber":"123456789", "countryCode":"ES"}`)

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))

		router.ServeHTTP(writer, request)
		json.NewDecoder(writer.Body).Decode(&response)
		twilioMock := app.TwilioClient.(*mocks.TwilioMock)

		assert.Equal(t, 1, twilioMock.TimesInvoked)
	})

	t.Run("POST /users does not create a new user if no phone or phone prefix is provided", func(t *testing.T) {
		ResetDatabase()

		_, router := SetupRouter()

		var noPrefixPayload = []byte(`{"phoneNumber":"123456789", "countryCode":"ES"}`)
		var noNumberPayload = []byte(`{"phonePrefix":"+34", "countryCode":"ES"}`)

		noPrefixWriter := httptest.NewRecorder()
		noNumberWriter := httptest.NewRecorder()
		noPrefixRequest, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(noPrefixPayload))
		noNumberRequest, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(noNumberPayload))

		router.ServeHTTP(noPrefixWriter, noPrefixRequest)
		router.ServeHTTP(noNumberWriter, noNumberRequest)

		assert.Equal(t, 422, noPrefixWriter.Code)
		assert.Equal(t, 422, noNumberWriter.Code)
		assert.Equal(t, true, DatabaseMissing("users_users", "phone_number='123456789' AND country_code='ES'"))
		assert.Equal(t, true, DatabaseMissing("users_users", "phone_prefix='+34' AND country_code='ES'"))
	})

	t.Run("POST /users does not create a new user if no country code is provided", func(t *testing.T) {
		ResetDatabase()

		_, router := SetupRouter()
		var payload = []byte(`{"phonePrefix":"+34", "phoneNumber":"123456789"}`)

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))

		router.ServeHTTP(writer, request)

		assert.Equal(t, 422, writer.Code)
		assert.Equal(t, true, DatabaseMissing("users_users", "phone_prefix='+34' AND phone_number='123456789'"))
	})

	t.Run("GET /users/:userId returns the requested user", func(t *testing.T) {
		ResetDatabase()

		user := factories.CreateUser(factories.UserFactoryInput{PhonePrefix: "34", PhoneNumber: "166666666", CountryCode: "ES"})

		var response ExpectedUserResponse
		_, router := SetupRouter()
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/users/1", nil)
		AuthenticateAs(1 /* just pass the raw ID and skip the DB call */, request)

		router.ServeHTTP(writer, request)
		json.NewDecoder(writer.Body).Decode(&response)

		assert.Equal(t, 200, writer.Code)
		assert.Equal(t, user.ID, response.ID)
		assert.Equal(t, user.Name, response.Name)
		assert.Equal(t, user.LastName, response.LastName)
		assert.Equal(t, user.PhonePrefix, response.PhonePrefix)
		assert.Equal(t, user.PhoneNumber, response.PhoneNumber)
		assert.Equal(t, user.CountryCode, response.CountryCode)
	})
}
