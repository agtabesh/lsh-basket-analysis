package main

import (
	"encoding/csv"
	"os"

	"github.com/agtabesh/lsh/types"
)

// readVectorsFromFile reads vectors (representing transactional data) from a CSV file
func readVectorsFromFile(path string) (map[string]types.Vector, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	datasetMap := make(map[string]types.Vector)
	// Iterate over the records (rows) in the CSV file, starting from index 1 (skipping header)
	for _, record := range records[1:] {
		userID := record[0]
		ItemName := record[2]
		if datasetMap[userID] == nil {
			datasetMap[userID] = make(types.Vector)
		}
		datasetMap[userID][types.VectorID(ItemName)] = 1
	}
	return datasetMap, nil
}
