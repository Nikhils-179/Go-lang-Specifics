// Write a function to implement a worker pool where several Goroutines process tasks concurrently
// The Pool should be able to distribute taks to workers and collect the results once they are done

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

type Customerss struct {
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

var (
	jobss = make(chan Customerss, 2)
	wg    sync.WaitGroup
)

func withConcurrency() {

	file, err := os.Open("customers-100.csv")
	if err != nil {
		log.Fatalf("Error openeing csv file %v", file)
	}
	defer file.Close()
	fmt.Println("Opened csv file....")
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reding records %v ", err)
	}

	num_workers := 5
	createWorkerPool(num_workers)

	go func() {
		for _, record := range records[1:] {
			index, err := strconv.Atoi(record[0])
			if err != nil {
				log.Fatalf("Error Occured %v", err)
			}
			customerss := Customerss{
				Index: index,
			}
			fmt.Println("Customer ", customerss.Index, " processing")
			jobss <- customerss
		}
		close(jobss)
	}()

}

func createWorkerPool(num int) {
	for i := 0; i < num; i++ {
		wg.Add(1)
		go work(i)
	}
}

func work(i int) {
	for job := range jobss {
		processJobs(i, job)
	}
	wg.Done()
}

func processJobs(i int, jobs Customerss) {
	fmt.Println("Worker ", i, "is running")
	cust := jobs
	updated_customer := Customers{
		Index: cust.Index,
	}
	fmt.Println("job ran ", updated_customer)

}
