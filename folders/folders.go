package folders

import (
	"github.com/gofrs/uuid"
)

// Return all folders with the specified OrgID as a struct
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {

	// Unused variable declaration removed

	// Fetch all folders with the requested OrgID
	foldersOfOrg, err := FetchAllFoldersByOrgID(req.OrgID, req.DataSet)

	// Return immediately on error
	if err != nil {
		return nil, err
	}

	/*
		Section removed as &v1 references the address of v1 which remains constant as it is used in the loop
		This means that all the elements in fp are dereferenced to whatever the value of v1 is at the time
		When the loop finishes, the value of v1 is the last folder in foldersOfOrg, so the program will output the same folder repetitively

		for _, v := range foldersOfOrg {
			folders = append(folders, *v)
		}

		var fp []*Folder

		for _, v1 := range folders {
			fp = append(fp, &v1)            // Logic error: &v1 is the same for all folders
		}
	*/

	// Find length of foldersOfOrg
	count := len(foldersOfOrg)

	// Return resulting folders as struct
	res := &FetchFolderResponse{Folders: foldersOfOrg, Count: count}

	return res, nil
}

// Fetch all folders with the specified OrgID
func FetchAllFoldersByOrgID(orgID uuid.UUID, dataSet []*Folder) ([]*Folder, error) {
	var folders []*Folder
	var err error

	if dataSet == nil {
		// Fetch folders from sample file
		folders, err = GetJSONData("sample.json")
		if err != nil {
			return nil, err
		}
	} else {
		folders = dataSet
	}

	// Iterate over all folders and append folders with specified OrgID to resFolders
	resFolders := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolders = append(resFolders, folder)
		}
	}

	return resFolders, nil
}
