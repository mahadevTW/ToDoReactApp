package handlers

import (
	"net/http"
	"github.com/gorilla/csrf"
	"encoding/json"
)

type token struct{
	CSRFToken string `json:csrftoken`
}

func CSRFHandler() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

		csrfToken := &token{CSRFToken:csrf.Token(r)}
		response, err := json.Marshal(csrfToken)

		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(response)

		return
	}
}
