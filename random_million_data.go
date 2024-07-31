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
	// Creating the zip file
	zipFile, err := os.Create("data.zip")
	if err != nil {
		fmt.Println("Error creating zip file:", err)
		return
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Creating a CSV file inside the zip
	csvFile, err := zipWriter.Create("Million_data.csv")
	if err != nil {
		fmt.Println("Error creating CSV file inside zip:", err)
		return
	}

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush() // Use Flush here because we write to a zip

	rand.Seed(time.Now().Unix())

	// Generating and writing 1 million random numbers
	for i := 0; i < 1000000; i++ {
		randomNumber := generateRandomNumber()
		if err := csvWriter.Write([]string{randomNumber}); err != nil {
			fmt.Println("Error writing to CSV:", err)
			return
		}
	}

	// Corrected the print statement's placement
	fmt.Println("1 million records written successfully to Million_data.csv inside data.zip")
}

func generateRandomNumber() string {
	firstDigit := rand.Intn(8) + 2                        // Ensures the first digit is not zero
	remainingDigits := rand.Int63n(900000000) + 100000000 // Generates the remaining 8 digits
	return strconv.Itoa(firstDigit) + strconv.FormatInt(remainingDigits, 10)
}
