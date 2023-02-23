package main

import (
	"encoding/csv"
	"os"
	"testing"
)

// Define a struct for a person with values and types
type Person struct {
	Name string
	Age  int
	Job  string
	City string
}

// TestWriteCSV tests writing persons to a CSV file
func TestWriteCSV(t *testing.T) {
	// Define sample data
	persons := []Person{
		{Name: "Ahmet Ünal", Age: 40, Job: "DevOps Lead Engineer", City: "Stuttgart"},
		{Name: "Anna Schmidt", Age: 34, Job: "Data Scientist", City: "München"},
		{Name: "Hans Müller", Age: 35, Job: "General Manager", City: "Berlin"},
		{Name: "Ursula Laien", Age: 55, Job: "Security Specialist", City: "Frankfurt"},
	}

	// Define expected output
	expected := [][]string{
		{"Ahmet Ünal", "DevOps Lead Engineer", "Stuttgart"},
		{"Anna Schmidt", "Data Scientist", "München"},
		{"Hans Müller", "General Manager", "Berlin"},
		{"Ursula Laien", "Security Specialist", "Frankfurt"},
	}

	// Create a temporary file for writing data
	file, err := os.CreateTemp("", "test-*.csv")
	if err != nil {
		t.Fatal("Failed to create temporary file")
	}
	defer os.Remove(file.Name())

	// Create a CSV writer object for writing data to file
	writer := csv.NewWriter(file)

	// Write each person record to the CSV file
	for _, person := range persons {
		record := []string{person.Name, person.Job, person.City}
		if err := writer.Write(record); err != nil {
			t.Fatalf("Failed to write record to file: %v", err)
		}
	}
	writer.Flush()

	// Read the data from the temporary file to verify it was written correctly
	file, err = os.Open(file.Name())
	if err != nil {
		t.Fatalf("Failed to open temporary file: %v", err)
	}
	defer file.Close()

	// Create a CSV reader object for reading data from file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	// Check if the records match the expected output
	for i, record := range records {
		for j, value := range record {
			if value != expected[i][j] {
				t.Fatalf("Record %d, column %d: expected %q, got %q", i, j, expected[i][j], value)
			}
		}
	}
}


