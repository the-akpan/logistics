package controllers

import (
	"encoding/json"
	"net/http"
	"tracka/pkg/config"
	"tracka/pkg/database"
	"tracka/pkg/utils"
)

func AuthSignin(res http.ResponseWriter, req *http.Request) {
	response := utils.CreateResponse(res)

	defer json.NewEncoder(res).Encode(response)

	var formStruct struct {
		Email, Password string
	}

	if !utils.DecodeRequest(&formStruct, response, res, req) {
		return
	}

	data := make(map[string]interface{})

	switch email := formStruct.Email; {
	case email == "":
		data["email"] = fieldRequired
	case !emailRX.MatchString(email):
		data["email"] = "Invalid email"
	}

	if formStruct.Password == "" {
		data["password"] = fieldRequired
	}

	if len(data) > 0 {
		response.Message = "Bad Request"
		response.Error = data
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := database.UserColl().GetUser(formStruct.Email)
	if err != nil || !user.CheckPassword(formStruct.Password) {
		response.Message = "Invalid credentials"
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := utils.CookieClaim{Email: user.Email, UserAgent: req.UserAgent(), IP: req.RemoteAddr}
	cookie, err := utils.GenerateCookie(&claims)
	if err != nil {
		response.Message = "Something went wrong"
		res.WriteHeader(http.StatusInternalServerError)
	}

	http.SetCookie(res, cookie)

	response.Message = "Logged in successfully"
	response.Data = user
	res.WriteHeader(http.StatusOK)
}

func AuthLogout(res http.ResponseWriter, req *http.Request) {
	response := utils.CreateResponse(res)

	defer json.NewEncoder(res).Encode(response)

	cookie, _ := utils.GenerateCookie(nil)
	http.SetCookie(res, cookie)

	res.WriteHeader(http.StatusResetContent)
	response.Message = "Logged out successfully"
}

func AuthResetPassword(res http.ResponseWriter, req *http.Request) {
	response := utils.CreateResponse(res)

	defer json.NewEncoder(res).Encode(response)

	var formStruct struct {
		Email string
	}

	if !utils.DecodeRequest(&formStruct, response, res, req) {
		return
	}

	data := make(map[string]interface{})

	switch email := formStruct.Email; {
	case email == "":
		data["email"] = fieldRequired
	case !emailRX.MatchString(email):
		data["email"] = "Invalid email"
	}

	response.Message = "Password reset to default if user exists"
	res.WriteHeader(http.StatusOK)

	user, err := database.UserColl().GetUser(formStruct.Email)
	if err != nil {
		return
	}

	defaultPasswod := config.Get().Admin.Password
	user.MakePassword(defaultPasswod)

	err = database.UserColl().UpdateUser(*user)
	if err != nil {
		response.Message = "Something went wrong"
		res.WriteHeader(http.StatusInternalServerError)
	}
}
