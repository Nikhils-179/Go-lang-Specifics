package main

import "fmt"

// ILogger interface
type ILogger interface {
	Log(message string)
}

// FileLogger struct
type FileLogger struct{}

// Log method for FileLogger
func (f FileLogger) Log(message string) {
	fmt.Println("FileLogger:", message)
}

// DatabaseLogger struct
type DatabaseLogger struct{}

// Log method for DatabaseLogger
func (d DatabaseLogger) Log(message string) {
	fmt.Println("DatabaseLogger:", message)
}

// DataAccessLayer struct
type DataAccessLayer struct {
	logger ILogger
}

// NewDataAccessLayer constructor
func NewDataAccessLayer(logger ILogger) *DataAccessLayer {
	return &DataAccessLayer{logger: logger}
}

// AddCustomer method
func (d *DataAccessLayer) AddCustomer(name string) {
	// Code to add customer to the database
	d.logger.Log("Customer added: " + name)
}

func main() {
	fileLogger := FileLogger{}
	dal1 := NewDataAccessLayer(fileLogger)
	dal1.AddCustomer("John Doe") // Output: FileLogger: Customer added: John Doe

	databaseLogger := DatabaseLogger{}
	dal2 := NewDataAccessLayer(databaseLogger)
	dal2.AddCustomer("Jane Doe") // Output: DatabaseLogger: Customer added: Jane Doe
}
