package main

import (
	"errors"
	"fmt"
	"flag")


type Args struct {
	date string
	car string
	price float64
	gallons float64
}


func SetupCmdArgs() Args {
	date_filled := flag.String("date", "0", "The date you filled up")
	car_filled :=  flag.String("car", "0", "The model of car filled")
	total_price :=  flag.Float64("price", 0, "Total price of gas")
	gallons_filled :=  flag.Float64("gallons", 0, "Amount of gallons filled")

	flag.Parse()

	return Args{
		date: *date_filled,
		car: *car_filled,
		price: *total_price,
		gallons: *gallons_filled,
	}
}

func (args *Args) validate_args() (error){
	if args.date == "0" {
		return errors.New("Date needs to be specified!")
	}

	if args.car == "0" {
		return errors.New("Car model needs to be specified!")
	}

	if args.price == 0 {
		return errors.New("Price of gas needs to be specified!")
	}

	if args.gallons == 0 {
		return errors.New("Total gallons needs to be specified!")
	}
	return nil
}

func main() {
	// check if db file exists

	// create db

	args := SetupCmdArgs()

	// Validate args
	err := args.validate_args()

	if err != nil {
		fmt.Println(err)
	}
	// Write to sqlite db
}
