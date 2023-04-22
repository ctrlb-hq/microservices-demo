package main

import (
	"dev0-hq/microservices-demo/go-svc/db"
	"fmt"
	"io/ioutil"
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
	router.HandleFunc("/pingPython", checkPythonStatus).Methods("GET")

	DBObj = db.SetupDB(dbHost, dbUserName, dbPassowrd, dbName, dbPort)

	fmt.Println("Server at " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func fetchNumber(w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("uuid")
	var number db.Number
	DBObj.First(&number, "uuid = ?", uuid)

	magicNumber := &MagicNumber{
		UUID:   uuid,
		Number: int32(number.Number) * 2,
	}

	json.NewEncoder(w).Encode(magicNumber)
}

func check(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong!"))
}

func checkPythonStatus(w http.ResponseWriter, r *http.Request) {
	pythonSvcHost := getEnv("PYTHON_SERVICE_HOST", "0.0.0.0")
	pythonSvcPort := getEnv("PYTHON_SERVICE_PORT", "30000")

	pythonEndpoint := fmt.Sprintf("http://%s:%s/ping", pythonSvcHost, pythonSvcPort)
	resp, err := http.Get(pythonEndpoint)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	w.Write(body)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
