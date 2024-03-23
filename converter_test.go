package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestConverterBehavesAsExpected(t *testing.T) {
	a := assert.New(t)

	csvFilePath := "./mockcsv_data.csv"
	jsonFilePath := "./mockjson_data.json"

	Convert(csvFilePath, jsonFilePath)

	jsonFile, _ := os.Open(jsonFilePath)
	byteResult, _ := io.ReadAll(jsonFile)

	var songs []Song
	json.Unmarshal(byteResult, &songs)

	fmt.Println(songs)

	a.Equal(3, len(songs), "3 songs found in converted file")

	_ = os.Remove(jsonFilePath)
}
