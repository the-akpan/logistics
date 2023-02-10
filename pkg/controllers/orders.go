package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tracka/pkg/database"
	"tracka/pkg/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

func OrdersList(res http.ResponseWriter, req *http.Request) {
	response := utils.CreateResponse(res)

	defer json.NewEncoder(res).Encode(response)

	var page int64 = 1
	var limit int64 = 20

	if req.URL.Query().Has("page") {
		tmp_page, err := strconv.Atoi(req.URL.Query().Get("page"))
		if err == nil {
			page = int64(tmp_page)
		}

	}
	if req.URL.Query().Has("limit") {
		tmp_limit, err := strconv.Atoi(req.URL.Query().Get("limit"))
		if err == nil {
			limit = int64(tmp_limit)
		}
	}

	orders, err := database.OrderColl().GetOrders(page, limit)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			response.Message = "Something went wrong" + err.Error()
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	response.Data = orders
}
