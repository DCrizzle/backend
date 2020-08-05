// +build mock

package main

import (
	"fmt"
	"net/http"
)

func usersHandler(secret, token, url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"message": "success"}`)
	}
}
