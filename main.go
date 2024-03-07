package main

// The code starts by importing necessary packages.
// `context`, `csv`, `os`, and `fmt` are standard Go packages.
// `sort` is used for sorting slices.
// `github.com/agtabesh/lsh` contains the Locality-Sensitive Hashing (LSH) implementation, and `github.com/agtabesh/lsh/types` contains custom types used in the LSH implementation.
import (
	"context"
	"fmt"

	"github.com/agtabesh/lsh"
	"github.com/agtabesh/lsh/types"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	// The `run` function starts by reading vectors (representing transactional data) from a CSV file named "Groceries_dataset.csv" using the `readVectorsFromFile` function.
	// It handles any errors that occur during file reading and returns an error if encountered.
	datasetMap, err := readVectorsFromFile("Groceries_dataset.csv")
	if err != nil {
		return err
	}

	// Defines the configuration for the Locality-Sensitive Hashing (LSH) algorithm.
	// The SignatureSize and BandSize parameters determine the size of the hash signature and the number of hash bands, respectively.
	config := lsh.LSHConfig{
		SignatureSize: 100,
		BandSize:      50,
	}

	// The code initializes components required for LSH, including the hash family (`hashFamily`), similarity measure (`similarityMeasure`), and storage mechanism (`store`).
	// In this case, `XXHASH64HashFamily` is used for hashing, `HammingSimilarity` for measuring similarity, and an in-memory store.
	hashFamily := lsh.NewXXHASH64HashFamily(config.SignatureSize)
	similarityMeasure := lsh.NewHammingSimilarity()
	store := lsh.NewInMemoryStore()

	// An LSH instance is created using the previously defined configuration, hash family, similarity measure, and store.
	instance, err := lsh.NewLSH(config, hashFamily, similarityMeasure, store)
	if err != nil {
		return err
	}

	// The code iterates over the dataset map obtained from the CSV file, and adds them to the LSH instance.
	ctx := context.Background()
	for i, vector := range datasetMap {
		vectorID := types.VectorID(i)
		err := instance.Add(ctx, vectorID, vector)
		if err != nil {
			return err
		}
	}

	// A sample vector is defined, representing a product ("white bread"), and the `QueryByVector` method is called to find similar vectors in the LSH instance.
	// The number of similar vectors to retrieve is specified by the `count` variable.
	vector := types.Vector{"white bread": 1}
	count := 100
	similarVectorsID, err := instance.QueryByVector(ctx, vector, count)
	if err != nil {
		return err
	}

	// The code iterates over the IDs of similar vectors retrieved and aggregates the associated items from the dataset map, counting their occurrences.
	// The top 10 associated items are extracted from the aggregated data, and the result is printed to the console.
	items := make(Items)
	for _, vectorID := range similarVectorsID {
		for item := range datasetMap[vectorID.String()] {
			items[item.String()]++
		}
	}
	n := 10
	result := items.Top(n)
	fmt.Println("Result:", result)

	return nil
}
