package main

import (
	"dev0-hq/microservices-demo/go-svc/db"
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type MagicNumber struct {
	UUID   string `json:"uuid"`
	Number int32  `json:"number"`
}

var DBObj *gorm.DB

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/fetchNumber", fetchNumber).Methods("GET")
	router.HandleFunc("/ping", check).Methods("GET")

	DBObj = db.SetupDB()

	fmt.Println("Server at 3099")
	log.Fatal(http.ListenAndServe(":3099", router))
}

func fetchNumber(w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("uuid")
	var number db.Number
	DBObj.First(&number, "uuid = ?", uuid)

	magicNumber := &MagicNumber{
		UUID:   uuid,
		Number: int32(number.Number),
	}

	json.NewEncoder(w).Encode(magicNumber)
}

func check(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong!"))
}
