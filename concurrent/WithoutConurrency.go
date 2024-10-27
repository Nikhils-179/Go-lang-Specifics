
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Customers struct {
	Index             int    `json:"index"`
	Customer_Id       int    `json:"customer_id"`
	First_Name        string `json:"firstname"`
	Second_Name       string `json:"secondname"`
	Last_Name         string `json:"lastname"`
	Company           string `json:"company"`
	City              string `json:"city"`
	Country           string `json:"country"`
	Phone1            int    `json:"phone1"`
	Phone2            int    `json:"phone2"`
	Email             string `json:"email"`
	Subscription_Date string `json:"subscription_date"`
	Website           string `json:"website"`
}

func withoutConcurrency(){
	start := time.Now()
	file, err := os.Open("customers-100.csv")
	if err != nil {
		log.Fatal("Error opening csv file", err)
	}
	defer file.Close()
	fmt.Println("Opened csv file....")

	fmt.Println("Starting reading file....")
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error occured while reading the file")
	}

	for _ , record := range records[1:] {
		index , err :=  strconv.Atoi(record[0])
		if err!=nil{
			log.Fatalf("Error converting string to integer %s",err)
		}
		Customers := Customers{
			Index: index,
		}

		processCustomer(Customers)
	}

	d := time.Since(start)
	fmt.Println(d)
}

func processCustomer(customer Customers ){
	fmt.Println("Starting process on customer " , customer.Index)
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Customer" , customer.Index  , "Data has been processed")
}