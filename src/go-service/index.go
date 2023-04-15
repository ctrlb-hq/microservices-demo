package main

import (
	"dev0-hq/microservices-demo/go-svc/db"
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"os"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type MagicNumber struct {
	UUID   string `json:"uuid"`
	Number int32  `json:"number"`
}

var DBObj *gorm.DB

func main() {
	port := getEnv("GO_SERVICE_PORT", "30002")
	dbHost := getEnv("DB_SERVICE_HOST", "0.0.0.0")
	dbUserName := getEnv("POSTGRES_USER", "postgres")
	dbPassowrd := getEnv("POSTGRES_PASSWORD", "mysecretpassword")
	dbName := getEnv("DB_SERVICE_NAME", "numbers")
	dbPort := getEnv("DB_SERVICE_PORT", "5432")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/fetchNumber", fetchNumber).Methods("GET")
	router.HandleFunc("/ping", check).Methods("GET")

	DBObj = db.SetupDB(dbHost, dbUserName, dbPassowrd, dbName, dbPort)

	fmt.Println("Server at "+port)
	log.Fatal(http.ListenAndServe(":"+port, router))
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

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
