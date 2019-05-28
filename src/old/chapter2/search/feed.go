package search

import (
	"encoding/json" // used to encode, decode json.
	"os"            // os functionalities like read write etc.
)

const dataFile = "data/data.json" // relative path to the data. GO can determine the file type from extension

// the above data has 3 fields, called name, uri, and type. We build a struct that matches the data feed. We need to create a struct that matches our type to use that in the program.
// in this case we are including the metadata 'json' to indicate what decoding needs to be used to create a type Feed.
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// In search.go we are expecting 2 parameters to be returned from this function. Actual feed and an error
// The above struct is used to create a slice of type Feed. and error as well.
// Data is unmarshelled in this function
func RetrieveFeeds() ([]*Feed, error) {
	// open dataFile
	// relative path must be provided to open the file.
	// this call will return 2 values, a pointer to the value of type File, and error which determines if call to open the file was successful or not.
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// close the file after the funciton returns
	defer file.Close()

	// Decode the file into a slice of pointers to Feed
	// Nil slice that contains pointers to Feed type
	var feeds []*Feed

	// decode the value returned from NewDecoder.
	// NewDecoder returns a pointer to a value of type Decoder. We can pass the decode mothod to that value with the address to the slice (&feed).
	// Decode method decodes file and populates the feed slice with type Feed values
	// Decode method accepts type interface{}.
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
