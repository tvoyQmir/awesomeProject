package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

var db *sql.DB

func rollHandler(w http.ResponseWriter, r *http.Request){
	log.Print("rollHandler(", w, r, ")")

	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_list.html")

		if err != nil {
			log.Fatal(err)
		}

		cars, err := dbGetCars()
		if err != nil {
			log.Fatal(err)
		}

		t.Execute(w, cars)
	}
}

func addCarManager(w http.ResponseWriter, r *http.Request) {
	log.Print("addCarManager(", w, r, ")")

	if r.Method == "GET" {
		var t, err = template.ParseFiles("simple_form.html")

		if err != nil {
			log.Fatal(err)
		}

		t.Execute(w, nil)
	} else {
		r.ParseForm()

		brand :=         r.Form.Get("brand")
		model :=         r.Form.Get("model")
		typeOfCarBody := r.Form.Get("typeOfCarBody")
		price, errPrice := strconv.Atoi(r.Form.Get("price"))

		err := dbAddNewCar(brand, model, typeOfCarBody, price)

		if err != nil || errPrice != nil {
			log.Fatal(err)
		}
	}
}

func GetPort() string {
	log.Print("GetPort")

	var port = os.Getenv("PORT")

	if port == "" {
		port = "4747"
		fmt.Println(port)
	}

	return ":" + port
}

func main() {
	log.Print("Starting...")
	var err = dbConnect()

	if err != nil {
		log.Fatal(err)
	}

	log.Print("HandleFunc")
	http.HandleFunc("/", rollHandler)
	http.HandleFunc("/add", addCarManager)
	log.Fatal(http.ListenAndServe(GetPort(), nil))
}