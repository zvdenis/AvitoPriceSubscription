package main

import (
	"awesomeProject/AvitoWatcher"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	d, err := sql.Open("mysql", "root:root@/avito_watch")
	if err != nil {
		panic(err)
	}
	defer d.Close()

	var subManager AvitoWatcher.SubscriptionManager
	subManager.DB = d

	var handler AvitoWatcher.Handler
	handler.SubManager = subManager

	var watcher AvitoWatcher.Watcher
	watcher.SubManager = subManager

	go watcher.WatchAvito()
	router := mux.NewRouter()
	router.HandleFunc("/", handler.Index).Methods("GET")
	router.HandleFunc("/", handler.Subscribe).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
