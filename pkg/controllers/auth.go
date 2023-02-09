package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"tracka/pkg/models"
	"tracka/pkg/utils"
)

func AuthSignin(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	response := Response{}

	defer json.NewEncoder(res).Encode(&response)

	var loginForm struct {
		Email, Password string
	}

	err := json.NewDecoder(req.Body).Decode(&loginForm)
	if err != nil {
		response.Message = "Error Decoding Values " + err.Error()
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	data := make(map[string]interface{})
	emailRX := regexp.MustCompile(`^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)

	switch email := loginForm.Email; {
	case email == "":
		data["email"] = fieldRequired
	case !emailRX.MatchString(email):
		data["email"] = "Invalid email"
	}

	if loginForm.Password == "" {
		data["password"] = fieldRequired
	}

	if len(data) > 0 {
		response.Message = "Bad Request"
		response.Error = data
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := models.UserColl().GetUser(loginForm.Email)
	if err != nil || !user.CheckPassword(loginForm.Password) {
		response.Message = "Invalid credentials"
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := utils.CookieClaim{Email: user.Email, UserAgent: req.UserAgent(), IP: req.RemoteAddr}
	cookie := utils.GenerateCookie(claims)

	http.SetCookie(res, cookie)

	response.Message = "Logged in successfully"
	response.Data = user
	res.WriteHeader(http.StatusOK)
}
