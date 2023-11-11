package folders

import (
	"encoding/json"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gofrs/uuid"
	"github.com/lucasepe/codename"
)

// These are all helper methods and fixed types.
// There's no real need for you to be editting these, but feel free to tweak it to suit your needs.
// If you do make changes here, be ready to discuss why these changes were made.

const DefaultOrgID = "c1556e17-b7c0-45a3-a6ae-9546248fb17a"
const RandomDataSetSize = -1 // acts as a flag to generate data of random size
const DefaultDataSetSize = 1000
const MinDataSetSize = 0
const MaxDataSetSize = 1000000

type Folder struct {
	// A unique identifier for the folder, must be a valid UUID.
	// For example: '00001d65-d336-485a-8331-7b53f37e8f51'
	Id uuid.UUID `json:"id"`
	// Name associated with folder.
	Name string `json:"name"`
	// The organisation that the folder belongs to.
	OrgId uuid.UUID `json:"org_id"`
	// Whether a folder has been marked as deleted or not.
	Deleted bool `json:"deleted"`
}

// Generate sample data of size 1000
// EDITED: generate sample data with specified dataSetSize, and also return expected number of occurrences of DefaultOrgID
func GenerateData(dataSetSize int) ([]*Folder, int, error) {
	var expectedCount int
	rng, err := codename.DefaultRNG() // RNG seed
	sampleData := []*Folder{}

	// ADDED: Randomise size if requested
	if dataSetSize == RandomDataSetSize {
		dataSetSize = rand.Intn(MaxDataSetSize)
	}

	for i := 1; i < dataSetSize; i++ {
		orgId := uuid.FromStringOrNil(DefaultOrgID)

		// Every 3 folders, have a random OrgID
		// EDITED: randomly distribute folders with unique / default OrgID
		if rand.Int()%2 == 1 {
			orgId = uuid.Must(uuid.NewV4())
		} else {
			expectedCount++ // Increment expectedCount for each folder with DefaultOrgID
		}

		// Randomly assign deleted status (using even/odd random integers)
		deleted := rand.Int() % 2

		sampleData = append(sampleData, &Folder{
			Id:      uuid.Must(uuid.NewV4()),   // random ID
			Name:    codename.Generate(rng, 0), // random 'hero-like' name
			OrgId:   orgId,
			Deleted: deleted != 0, // converting integer to Boolean
		})
	}

	return sampleData, expectedCount, err
}

// Applies indentation to JSON output and prints it
// EDITED: just prettify it to allow code reuse when testing
func Prettify(rawJSON interface{}) string {
	prettifiedJSON, _ := json.MarshalIndent(rawJSON, "", "\t") // applying indentation
	return string(prettifiedJSON)
}

// Fetch folder structs from sample file
func GetSampleData() ([]*Folder, error) {
	// Get absolute path of this static.go file
	_, filename, _, _ := runtime.Caller(0)
	// fmt.Println("Script file:", filename)          // No need to expose directory structure

	// Get absolute path of the sample.json file in the same directory
	basePath := filepath.Dir(filename)
	sampleFilePath := filepath.Join(basePath, "sample.json")
	// fmt.Println("Sample file:", sampleFilePath)    // No need to expose directory structure

	// Open sample.json file
	sampleFile, err := os.Open(sampleFilePath)
	if err != nil {
		return nil, err // Return error instead of panicking
	}
	defer sampleFile.Close() // File will close once it is read below

	// Read and return sample.json as a slice of (pointers to) folders
	jsonByte, err := io.ReadAll(sampleFile)
	if err != nil {
		return nil, err
	}
	folders := []*Folder{}
	json.Unmarshal(jsonByte, &folders)
	return folders, nil
}
