package folders

import "github.com/gofrs/uuid"

// EDITED: also allows for dataset other than sample.json to be supplied
type FetchFolderRequest struct {
	OrgID   uuid.UUID
	DataSet []*Folder
}

// EDITED: also contains number of folders returned
type FetchFolderResponse struct {
	Folders []*Folder
	Count   int
}
