package main

import
(
	"database/sql"
	"fmt"
	"log"
)

type Car struct{
	Brand string
	Model string
	TypeOfCarBody string
	Price int
}

const (
	USERNAME = "postgres"
	PASS     = "user"
	DBNAME   = "db"
)

func dbConnect() error {
	log.Print("dbConnect()")

	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", USERNAME, PASS, DBNAME))

	if err != nil {
		return err
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS CARS (brand text, model text, typeOfCarBody text, price integer )"); err != nil {
		log.Fatal("Error. Couldn't create table: ", err)
		return err
	}

	return nil
}

func dbAddNewCar(brand string, model string, typeOfCarBody string, price int) error {
	log.Print("dbAddNewCar(", brand, ", ", model, ", ", typeOfCarBody, ", ", price, ")")

	var query = "INSERT INTO CARS VALUES ($1, $2, $3, $4)"

	var _, err = db.Exec(query, brand, model, typeOfCarBody, price)

	if err != nil {
		log.Fatal("Error. Couldn't insert into table: ", err)
		return err
	}

	return nil
}

func dbGetCars() ([]Car, error) {

	log.Print("dbGetCars()")
	var cars []Car

	var query, err = db.Prepare("SELECT * FROM CARS")
	if err != nil {
		log.Fatal("Error. Couldn't select from table")
		return cars, err
	}

	result, err := query.Query()
	if err != nil {
		log.Fatal("Error. Couldn't execute query: ", err)
		return cars, err
	}

	var car Car
	for result.Next() {
		err = result.Scan(&car.Brand, &car.Model, &car.TypeOfCarBody, &car.Price)

		if err != nil {
			log.Fatal("Error. Couldn't scan: ", err)
			return cars, err
		}

		log.Print("append: brand= ", car.Brand, ", model= ", car.Model, ", type= ", car.TypeOfCarBody, ", price= ", car.Price)
		cars = append(cars, car)
	}

	return cars, err
}