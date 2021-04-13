package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rGustave07/fbhcapi/api/model"
	"github.com/rGustave07/fbhcapi/api/model/apperrors"
	"github.com/rGustave07/fbhcapi/api/model/mockedmodels"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignup(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("Email and password Required", func(t *testing.T) {
		// We just want to show that it's not called in this case
		mockUserService := new(mockedmodels.MockUserService)
		mockUserService.On(
			"Signup",
			mock.AnythingOfType("*gin.Context"),
			mock.AnythingOfType("*model.User"),
		).Return(nil)

		// response recorder for getting written http response
		rr := httptest.NewRecorder()

		// Don't need middleware as we don't yet have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		// create a request body with empty email and pasword
		reqBody, err := json.Marshal(gin.H{
			"email":    "",
			"password": "",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertNotCalled(t, "Signup")
	})

	t.Run("Invalid email", func(t *testing.T) {
		// We just want this to show that it's not called in this case
		mockUserService := new(mockedmodels.MockUserService)
		mockUserService.On(
			"Signup", mock.AnythingOfType("*gin.Context"),
			mock.AnythingOfType("*model.User"),
		).Return(nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email":    "bob@bob",
			"password": "supersecret1234",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertNotCalled(t, "Signup")
	})

	t.Run("Password too short", func(t *testing.T) {
		// We just want this to show that it's not called in this case
		mockUserService := new(mockedmodels.MockUserService)
		mockUserService.On(
			"Signup", mock.AnythingOfType("*gin.Context"),
			mock.AnythingOfType("*model.User"),
		).Return(nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email":    "bob@blah.com",
			"password": "inval",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertNotCalled(t, "Signup")
	})

	t.Run("Password too long", func(t *testing.T) {
		// We just want this to show that it's not called in this case
		mockUserService := new(mockedmodels.MockUserService)
		mockUserService.On(
			"Signup", mock.AnythingOfType("*gin.Context"),
			mock.AnythingOfType("*model.User"),
		).Return(nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email":    "bob@blah.com",
			"password": "SuperLongInvalidPasswordThatGoesPastTheSetValidationInTheHandlerFunction",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertNotCalled(t, "Signup")
	})

	t.Run("Error calling UserService", func(t *testing.T) {
		u := &model.User{
			Email:    "blah@blah.com",
			Password: "avalidpassword",
		}

		mockUserService := new(mockedmodels.MockUserService)
		mockUserService.On(
			"Signup",
			mock.AnythingOfType("*gin.Context"),
			u,
		).Return(apperrors.NewConflictErr("User Already Exists", u.Email))

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email":    u.Email,
			"password": u.Password,
		})
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, 409, rr.Code)
		mockUserService.AssertExpectations(t)
	})

	t.Run("Successful Token Creation", func(t *testing.T) {
		u := &model.User{
			Email:    "test@mctester.com",
			Password: "avalidpassword",
		}

		mockTokenResponse := &model.TokenPair{
			IDToken:      "idToken",
			RefreshToken: "refreshToken",
		}

		mockUserService := new(mockedmodels.MockUserService)
		mockTokenService := new(mockedmodels.MockTokenService)

		mockUserService.
			On("Signup", mock.AnythingOfType("*gin.Context"), u).
			Return(nil)

		mockTokenService.
			On("NewPairFromUser", mock.AnythingOfType("*gin.Context"), u, "").
			Return(mockTokenResponse, nil)

		// A response recorder to track http response
		rr := httptest.NewRecorder()

		// Don't need middleware as we don't have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:            router,
			UserService:  mockUserService,
			TokenService: mockTokenService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email":    u.Email,
			"password": u.Password,
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"tokens": mockTokenResponse,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())

		mockUserService.AssertExpectations(t)
		mockTokenService.AssertExpectations(t)
	})

	t.Run("Failed Token Creation", func(t *testing.T) {
		u := &model.User{
			Email:    "test@mctester.com",
			Password: "avalidpassword",
		}

		mockErrorResponse := apperrors.NewInternalErr()

		mockUserService := new(mockedmodels.MockUserService)
		mockTokenService := new(mockedmodels.MockTokenService)

		mockUserService.
			On("Signup", mock.AnythingOfType("*gin.Context"), u).
			Return(nil)
		mockTokenService.
			On("NewPairFromUser", mock.AnythingOfType("*gin.Context"), u, "").
			Return(nil, mockErrorResponse)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:            router,
			UserService:  mockUserService,
			TokenService: mockTokenService,
		})

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email":    u.Email,
			"password": u.Password,
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"error": mockErrorResponse,
		})
		assert.NoError(t, err)

		assert.Equal(t, mockErrorResponse.Status(), rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())

		mockUserService.AssertExpectations(t)
		mockTokenService.AssertExpectations(t)
	})
}
