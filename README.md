# CSV To JSON Converter

A converter tool which takes a supplied CSV file containing song data from a Myriad instance, and returns a JSON file with just song title, artist and playout ID . 

To run: 

Clone the project 

```
$ git clone git@github.com:shreypuranik/csv-to-json-converter.git
```

Save the Myriad extracted CSV file in the project root as `csv_data.csv`. 

Run via cli 

```
$ go run .        
** STARTING CONVERT**
JSON file now available at ./data.json
** COMPLETED CONVERT**

```

Once completed, you should see `data.json` in the project root which contains song title, artist and ID for each row in the CSV file. 