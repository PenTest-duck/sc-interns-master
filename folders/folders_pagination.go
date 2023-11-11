package folders

//"github.com/gofrs/uuid"

const PageSize = 10

// Return all folders with the specified OrgID as a struct
func GetPaginatedFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {

	// Fetch all folders with the requested OrgID in sample.json
	foldersOfOrg, err := FetchAllFoldersByOrgID(req.OrgID, nil)

	// Return resulting folders as struct
	var res *FetchFolderResponse
	res = &FetchFolderResponse{Folders: foldersOfOrg}

	return res, err
}

func GetNextChunk(token string) {

}

// Ensuring unique tokens; IDOR?
