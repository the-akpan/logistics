package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"tracka/pkg/database"
	"tracka/pkg/models"
	"tracka/pkg/utils"

	"github.com/biter777/countries"
	"github.com/gorilla/mux"
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
		if err != mongo.ErrNoDocuments {
			response.Message = "Something went wrong" + err.Error()
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	response.Data = orders
}

func OrdersCreate(res http.ResponseWriter, req *http.Request) {
	response := utils.CreateResponse(res)
	defer json.NewEncoder(res).Encode(response)

	formStruct := models.OrderIn{}

	if !utils.DecodeRequest(&formStruct, response, res, req) {
		return
	}

	data := make(map[string]interface{})

	formStruct.From = strings.TrimSpace(formStruct.From)
	if formStruct.From == "" {
		data["from"] = fieldRequired
	} else {
		if country := countries.ByName(formStruct.From); country == countries.Unknown {
			data["from"] = fmt.Sprintf("Invalid country `%s`", formStruct.From)
		} else {
			formStruct.From = country.String()
		}
	}

	formStruct.To = strings.TrimSpace(formStruct.To)
	if formStruct.To == "" {
		data["to"] = fieldRequired
	} else {
		if country := countries.ByName(formStruct.To); country == countries.Unknown {
			data["to"] = fmt.Sprintf("Invalid country `%s`", formStruct.To)
		} else {
			formStruct.To = country.String()
		}
	}

	formStruct.Mode = strings.TrimSpace(formStruct.Mode)
	if formStruct.Mode == "" {
		data["mode"] = fieldRequired
	} else {
		if res := utils.ContainsString(database.MODES, formStruct.Mode); res == "" {
			data["mode"] = fmt.Sprintf("Invalid transport mode %s", formStruct.Mode)
		} else {
			formStruct.Mode = res
		}
	}

	if formStruct.CreatedAt.Equal(time.Time{}) {
		formStruct.CreatedAt = time.Now().UTC()
	}

	if len(data) > 0 {
		response.Message = "Bad Request"
		response.Error = data
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	order, err := database.OrderColl().CreateOrder(&formStruct)
	if err != nil {
		response.Message = "Something went wrong"
		res.WriteHeader(http.StatusInternalServerError)
	}

	response.Message = "Order Created"
	response.Data = order
	res.WriteHeader(http.StatusCreated)
}

func OrderGet(res http.ResponseWriter, req *http.Request) {
	response := utils.CreateResponse(res)
	defer json.NewEncoder(res).Encode(response)
	tracker := mux.Vars(req)["tracker"]

	order, err := database.OrderColl().GetOrder(tracker)
	if err != nil {
		response.Message = "Order not found"
		res.WriteHeader(http.StatusNotFound)
		return
	}

	if order.Updates == nil {
		order.Updates = make([]*models.OrderUpdate, 0)
	}

	response.Message = "Order found"
	response.Data = order
}
