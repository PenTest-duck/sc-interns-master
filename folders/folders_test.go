package folders

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/gofrs/uuid"
)

// Run all tests for sample.json
func Test_GetAllFoldersWithSampleJSON(t *testing.T) {

	const OrgIDForOneFolder = "2a14c480-3bef-4c53-ac1c-6c5fe44e54ef"
	const OrgIDForNoFolders = "a2f73842-85c1-4981-80c0-ea3c297538aa"
	const InvalidOrgID = "INVALID"

	// Table-driven unit tests
	var tests = []struct {
		id    int
		name  string
		orgID string
	}{
		{1, "Test correct 666 folders returned with default OrgID", DefaultOrgID},
		{2, "Test correct 1 folder returned for specified OrgID", OrgIDForOneFolder},
		{3, "Test 0 folders returned for random OrgID that does not match any folders", OrgIDForNoFolders},
		{4, "Test 0 folders returned for invalid UUID OrgID", InvalidOrgID},
	}

	// Iterate over the tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Fetch expected results from file
			JSONFilePath := "test-data/expected-test" + strconv.Itoa(test.id) + ".json"
			expectedResult, err := GetJSONData(JSONFilePath)
			if err != nil {
				t.Fatalf("Error reading test data: %v", err)
			}

			// Prepare request to fetch all folders with specified OrgID
			req := &FetchFolderRequest{
				OrgID: uuid.FromStringOrNil(test.orgID),
			}

			// Run the function and prettify output
			res, err := GetAllFolders(req)
			if err != nil { // Ensure no fatal errors
				t.Fatalf("Error with GetAllFolders(): %v", err)
			}

			// Ensure result is correct
			if !reflect.DeepEqual(res.Folders, expectedResult) {
				t.Errorf("Result for sample.json is incorrect")
			}
		})
	}
}

// Run all tests for randomly generated folders
func Test_GetAllFoldersWithRandomData(t *testing.T) {

	// Table-driven unit tests
	var tests = []struct {
		id          int
		name        string
		dataSetSize int
	}{
		{1, "Test with default dataset size of 1000", DefaultDataSetSize},
		{2, "Load test with max dataset size of 1000000", MaxDataSetSize},
		{3, "Test with min dataset size of 0", MinDataSetSize},
	}

	// Iterate over the tests
	for _, test := range tests {
		t.Run("Test with random data", func(t *testing.T) {
			// Generate random test data
			generatedData, expectedCount, err := GenerateData(test.dataSetSize)
			if err != nil {
				t.Fatalf("Error generating random test data: %v", err)
			}

			// Prepare request to fetch all folders with DefaultOrgID, supplying the random test data
			req := &FetchFolderRequest{
				OrgID:   uuid.FromStringOrNil(DefaultOrgID),
				DataSet: generatedData,
			}

			// Run the function
			res, err := GetAllFolders(req)
			if err != nil {
				t.Fatalf("Error with GetAllFolders(): %v", err)
			}

			// Check the number of occurrences of DefaultOrgID is as expected
			if res.Count != expectedCount {
				t.Errorf("Result for random dataset is incorrect")
			}
		})
	}
}
