package controller

import (
	"fmt"
	"net/http"
)

// GetReferees ...
func GetReferees() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("referee controller active")
	}
}

// AddReferee ...
func AddReferee() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("referee controller active")
	}
}
