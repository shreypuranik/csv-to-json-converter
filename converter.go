package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type Song struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

// Convert reads in the contents of csv_data.csv
// and generates a json file with song id, title and artist.
func Convert() {
	fmt.Println("** STARTING CONVERT**")

	jsonFileLocation := "./data.json"
	csvFile, err := os.Open("./csv_data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var songsArray []Song

	for index, individualSong := range csvData {
		if index == 0 {
			continue //workaround to avoid storing column headers from csv as song data
		}

		song := Song{
			ID:     individualSong[1],
			Title:  individualSong[2],
			Artist: individualSong[3],
		}

		songsArray = append(songsArray, song)
	}

	jsonData, err := json.Marshal(songsArray)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonFile, err := os.Create(jsonFileLocation)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()

	fmt.Printf(`JSON file now available at %s`, jsonFileLocation)
	fmt.Println()
	fmt.Println("** COMPLETED CONVERT**")
}
