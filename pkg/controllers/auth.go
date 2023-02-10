package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"tracka/pkg/config"
	"tracka/pkg/models"
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

	user, err := models.UserColl().GetUser(formStruct.Email)
	if err != nil || !user.CheckPassword(formStruct.Password) {
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

func AuthLogout(res http.ResponseWriter, req *http.Request) {
	log.Println("Hello")
	response := utils.CreateResponse(res)

	defer json.NewEncoder(res).Encode(response)

	res.WriteHeader(http.StatusNoContent)
	response.Message = "Logged out successfully"
}

func AuthRequestReset(res http.ResponseWriter, req *http.Request) {
	response := utils.CreateResponse(res)

	defer json.NewEncoder(res).Encode(response)

	var formStruct struct {
		Email string
	}

	if !utils.DecodeRequest(formStruct, response, res, req) {
		return
	}

	data := make(map[string]interface{})

	switch email := formStruct.Email; {
	case email == "":
		data["email"] = fieldRequired
	case !emailRX.MatchString(email):
		data["email"] = "Invalid email"
	}

	response.Message = "Reset email sent if user exists"
	res.WriteHeader(http.StatusOK)

	_, err := models.UserColl().GetUser(formStruct.Email)
	if err == nil {
		// Send email
		log.Println("Sending email to", formStruct.Email)
	}
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

	user, err := models.UserColl().GetUser(formStruct.Email)
	if err != nil {
		return
	}

	defaultPasswod := config.Get().Admin.Password
	user.MakePassword(defaultPasswod)

	err = models.UserColl().UpdateUser(*user)
	if err != nil {
		response.Message = "Something went wrong"
		res.WriteHeader(http.StatusInternalServerError)
	}
}
