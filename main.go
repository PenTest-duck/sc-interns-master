package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

func main() {
	var orgID string
	var paginate string

	// ADDED: allow for arbitrary OrgID input from command line
	fmt.Print("OrgID to search for (leave blank for default): ")
	fmt.Scanf("%v", &orgID)

	// Substitute default OrgID
	if orgID == "" {
		orgID = folders.DefaultOrgID
	}

	// Prepares a request containing the UUID of organisation to get folders for
	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(orgID),
	}

	// ADDED: ask whether to paginate result
	fmt.Print("Paginate result? (y/N): ")
	fmt.Scanf("%v", &paginate)

	// Fetches slice of folders with the requested OrgID
	res, err := folders.GetAllFolders(req)

	// Display error (if any)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if paginate == "y" || paginate == "Y" {
		pageSize := folders.DefaultPagesize
		pageNum := 1

		// Get size of each page until it is at least 1
		fmt.Print("Page size (leave blank for default): ")
		fmt.Scanf("%d", &pageSize)

		for pageSize < 1 {
			fmt.Println("Page size must be at least 1")

			fmt.Print("Page size (leave blank for default): ")
			fmt.Scanf("%d", &pageSize)
		}

		// Immediately print first page, then ask to print specified pages
		for pageNum != 0 {
			page := folders.GetPage(res, pageNum, pageSize)
			folders.PrettyPrint(page)

			fmt.Print("Page to print (0 to exit): ")
			fmt.Scanf("%d", &pageNum) // If non-digit entered, last stored value of pageNum used
		}

	} else {
		// If no pagination, just prettify and print all returned folders
		folders.PrettyPrint(res)
	}
}
