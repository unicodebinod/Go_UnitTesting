package main

import (
	"encoding/csv" // the csv package is used to read and write CSV files
	"fmt"          //the fmt package is used to print the output to the console.
	"os"           //the os package is used to interact with the operating system, such as opening and closing files
	"strconv"      //the strconv package is used to convert a string to an integer
     "testing"    //the testing package is used to run tests.

)

//function  generateTestData  generates test data in CSV format
//This code creates two variables: headers and data. 
//headers is a slice of strings that contains the names of the fields in the CSV file. 
//data is a slice of slices of strings that contains some example records with information about people.
func generateTestData() {
	headers := []string{"Name", "Nachname", "Alter", "Groesse", "Beruf", "Stadt", "Verheiratet"}
	data := [][]string{
	
		{"Anna", "Arti", "25", "165", "Ärztin", "München", "Nein"},
		{"Peter", "Meier", "45", "180", "Verkäufer", "Hamburg", "Ja"},
		{"Marie", "Schmidt", "28", "170", "Designerin", "Frankfurt", "Nein"},
		{"Hans", "Müller", "37", "185", "Lehrer", "Köln", "Ja"},
	}
 // this code creates a new file called "testdata.csv" using the os.Create function. 
 //if there is an error, it will cause a panic. 
 //the defer keyword ensures that the file is closed before the function returns.
	file, err := os.Create("testdata.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
// this code creates a new CSV writer using the csv.NewWriter function. 
//it then writes the headers and data slices to the file using the writer.Write method. 
//Finally, it flushes any buffered data using the writer.Flush method.
	writer := csv.NewWriter(file)
	writer.Write(headers)
	for _, row := range data {
		writer.Write(row)
	}
	writer.Flush()
}
//The BenchmarkReadCSV function is the main part of the program that measures
// the performance of reading the CSV data.
// It calls the generateTestData function to create a new CSV file 
//and then reads that file using the csv package.
func BenchmarkReadCSV(b *testing.B) {

//This line calls the generateTestData function to create a new CSV file.
	generateTestData()

	//This loop will execute the following code b.N times, 
	//where b.N is a value that is set by the testing framework 
	//to determine how many times the benchmark should be run.
	for i := 0; i < b.N; i++ {
		file, err := os.Open("testdata.csv") //The CSV file is opened using the os.Open function 
		if err != nil {
			panic(err)
		}
		defer file.Close()

		reader := csv.NewReader(file)
		var rows [][]string
		for {
			row, err := reader.Read()
			if err != nil {
				break
			}
			rows = append(rows, row)
		}
		// Ignore header row
		rows = rows[1:]

		for _, row := range rows {
			age, _ := strconv.Atoi(row[2])
			height, _ := strconv.Atoi(row[3])
			married := row[6] == "Ja"
			person := struct {
				Name       string
				LastName   string
				Age        int
				Height     int
				Profession string
				City       string
				Married    bool
			}{
				row[0], row[1], age, height, row[4], row[5], married,
			}
			// Do something with the person struct, e.g. print it
			fmt.Println(person)
		}
	}
}

func main() {
	// Run the benchmark
	testing.Benchmark(BenchmarkReadCSV)
}
