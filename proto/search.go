package proto


type PhotosSearch struct {
	Results []Photo `json:"results"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
}

type UsersSearch struct {
	Results []User `json:"results"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
}

type CollectionsSearch struct {
	Results []Collection `json:"results"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
}


