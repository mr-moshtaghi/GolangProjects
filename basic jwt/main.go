package main

import (
	"fmt"
	"net/http"
)

func GetJwt(w http.ResponseWriter, r *http.Request) {
	if r.Header["Access"] != nil {
		if r.Header["Access"][0] != api_key {
			return
		} else {
			token, err := CreateJWT()
			if err != nil {
				return
			}
			fmt.Fprint(w, token)
		}
	}
}
