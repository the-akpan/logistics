package routes

import (
	"encoding/json"
	"net/http"
	"strings"
	"tracka/pkg/config"
	"tracka/pkg/utils"
)

func RemoveSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		req.URL.Path = strings.TrimSuffix(req.URL.Path, "/")
		next.ServeHTTP(res, req)
	})
}

func authRequired(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		cookieGen := config.Get().Cookie
		cookie, err := req.Cookie(cookieGen.Name)
		if err == nil {
			value := make(map[string]interface{})
			if err = cookieGen.Cookie.Decode(cookie.Name, cookie.Value, &value); err == nil {
				next.ServeHTTP(res, req)
				return
			}
		}
		response := utils.CreateResponse(res)
		res.WriteHeader(http.StatusUnauthorized)
		response.Message = "Not Authorized"
		json.NewEncoder(res).Encode(response)
	})
}
