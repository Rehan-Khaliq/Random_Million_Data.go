package main

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	zipfile, err := os.Create("awais_data.zip") // Creation Of Zip File
	if err != nil {
		fmt.Println("Found An Error To Create Zip File:", err) // if file not created show this error
		return                                                 // return program if error found
	}
	defer zipfile.Close()
	zipwriter := zip.NewWriter(zipfile) // write data into zip file
	defer zipwriter.Close()             // After Writing Data file closed

	csvfile, err := zipwriter.Create("random_data.csv") // creation of csv file is file zip
	if err != nil {
		fmt.Println("Found Error To Create CSV File:", err) // if file not created show this error
		return                                              // Return Program If Error found
	}
	cswriter := csv.NewWriter(csvfile)
	defer cswriter.Flush() // All the data stored before exiting the program

	rand.Seed(time.Now().UnixNano()) // Used to generate Random number and it is very important

	for i := 0; i < 1000000; i++ {
		randomNumber := generateRandomNumber() //generateRandomNumber is builtin function to genrate random number
		if err := cswriter.Write([]string{randomNumber}); err != nil {
			fmt.Println("Found Error To Store String Data Into Csv file:", err)
			// write random number in string format into csv file
		}
	}
	fmt.Println("1 Million Data Successfully Stored In Csv File:")
	// Run this line outside the function if data created successfully
}
func generateRandomNumber() string {
	firstDigit := rand.Intn(8) + 2                  // shows number started from 2 ends on 9
	lastDigit := rand.Int63n(900000000) + 100000000 // generate random digit that always 8 digit long
	return strconv.Itoa(firstDigit) + strconv.FormatInt(lastDigit, 10)
}
