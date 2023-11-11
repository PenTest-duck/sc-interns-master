package folders

import "github.com/gofrs/uuid"

// EDITED: also allows for dataset other than sample.json to be supplied + arbitrary page size
type FetchFolderRequest struct {
	OrgID    uuid.UUID
	DataSet  []*Folder
	PageSize int
}

// EDITED: also contains number of folders returned (for added functionality)
type FetchFolderResponse struct {
	Folders []*Folder
	Count   int
}

// ADDED: returns folders within the page and tokens to other pages
// Token is 0 if there are no more pages in the direction (previous/next)
type PaginatedFetchFolderResponse struct {
	Folders      []*Folder
	TotalCount   int
	PageSize     int
	CurrentPage  int
	PreviousPage int
	NextPage     int
	FirstPage    int
	LastPage     int
}
