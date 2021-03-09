package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	model "fbhc.com/api/main/models/users"
	"gorm.io/gorm"
)

// UnmarshalUserJSON into golang struct
func UnmarshalUserJSON(incData []byte) (*model.User, error) {
	var data *model.User
	err := json.Unmarshal(incData, &data)

	return data, err
}

// AddNewUser does stuff
func AddNewUser(db *gorm.DB) http.HandlerFunc {
	// Maybe middleware here?
	return func(w http.ResponseWriter, r *http.Request) {

		request, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		userData, err := UnmarshalUserJSON(request)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userData.CreatedAt = time.Now()
		userData.UpdatedAt = time.Now()

		_, err = userData.HashOwnPassword()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		db.Select(
			"FirstName",
			"LastName",
			"Email",
			"Password",
			"CreatedAt",
			"UpdatedAt",
		).Create(userData)

		w.WriteHeader(http.StatusOK)
	}
}

// GetUsers ...
func GetUsers(db *gorm.DB) http.HandlerFunc {
	// Maybe middleware here?
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("userController activated")
	}
}
