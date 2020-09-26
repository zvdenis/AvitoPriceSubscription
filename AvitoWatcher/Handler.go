package AvitoWatcher

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Subscription struct {
	Url   string `json:"url"`
	Email string `json:"email"`
}

type Response struct {
	Message string `json:"message"`
}

//Отвечает за оброаботку запроосов
type Handler struct {
	SubManager SubscriptionManager
}

//Базовая страница
func (handler Handler) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}


//Обработка запроса подписки на объявление
func (handler Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var subscription Subscription
	json.NewDecoder(r.Body).Decode(&subscription)
	var response Response
	response.Message = "Success"
	fail := false

	price, id, err := GetPrice(subscription.Url)

	if !IsEmailValid(subscription.Email) {
		response.Message = "Invalid email"
		fail = true
	}

	if err != nil {
		response.Message = "Invalid url"
		fail = true
	}
	json.NewEncoder(w).Encode(response)

	if !fail {
		handler.SubManager.addSubscription(subscription, price, id)
	}
}
