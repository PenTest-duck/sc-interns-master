package folders

import "math"

func GetPage(folders *FetchFolderResponse, currentPage int, pageSize int) *PaginatedFetchFolderResponse {

	var previousPage int
	var nextPage int

	// Initialise empty slice of folders
	// This allows it to JSON.MarshalIndent() into "[]", instead of "null"
	paginatedFolders := []*Folder{}

	// Get number of folders and number of pages (totalCount ceiling divided by pageSize)
	totalCount := folders.Count
	lastPage := int(math.Ceil(float64(totalCount) / float64(pageSize)))

	// Ensure currentPage is within range
	if 1 <= currentPage && currentPage <= lastPage {

		// Calculate start index for page slice
		startIndex := pageSize * (currentPage - 1)

		// If on last page, fetch from startIndex to the end of slice
		// This prevents null padding in returned folders slice
		if currentPage == lastPage {
			paginatedFolders = folders.Folders[startIndex:]
		} else {
			// If not on last page, fetch page from startIndex to endIndex
			endIndex := pageSize * currentPage
			paginatedFolders = folders.Folders[startIndex:endIndex]
		}

		// Set previous and next page
		// If on last page, nextPage is left at 0
		previousPage = currentPage - 1
		if currentPage != lastPage {
			nextPage = currentPage + 1
		}
	}

	// Construct the response
	res := &PaginatedFetchFolderResponse{
		Folders:      paginatedFolders,
		TotalCount:   totalCount,
		PageSize:     pageSize,
		CurrentPage:  currentPage,
		PreviousPage: previousPage,
		NextPage:     nextPage,
		FirstPage:    1,
		LastPage:     lastPage,
	}

	return res
}

/*

RATIONALE

Instead of generating a random token that leads to the next page as shown in the example,
I decided to develop a more fully-featured pagination mechanism which stores a variety of
metadata allows the client to know the context around the pages they are viewing, including
total number of folders, page size, current/previous/next page, and first/last page.

This structuring was modelled on the Australian Government's API Design Standard at
https://api.gov.au/sections/pagination.html.

As per the standard, instead of erroring or returning a null object, the function will return []
if the specified page is out of range or there are no folders in the page.

The function has the unpaginated folders slice passed in, along with the target page and the
size of each page. It calculates the low and high bounds for the page and returns the slice of
folders within that page, along with updated metadata.

Arguably this method of pagination is slightly more prone to IDOR, but it allows the client to easily
traverse the pages and the server does not need to worry about generating (unique) tokens.

I had initially thought about generating a slice of slices, with each nested slice holding a page of
folders, then using indices to navigate them. However, I thought that would cause unnecessary
processing (especially for a large dataset), when a client might only need one page.

*/
