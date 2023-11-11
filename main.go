package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

func main() {

	// EDITED: allow for arbitrary OrgID input from command line
	var orgID string
	fmt.Print("OrgID to search for: ")
	fmt.Scanf("%v", &orgID)

	// Substitute default OrgID
	if orgID == "default" {
		orgID = folders.DefaultOrgID
	}

	// Prepares a request containing the UUID of organisation to get folders for
	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(orgID),
	}

	// Fetches slice of (pointers to) folders with the requested OrgID
	res, err := folders.GetAllFolders(req)

	// Display error (if any)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// Prettify and print returned folders
	fmt.Println(folders.Prettify(res))
}
