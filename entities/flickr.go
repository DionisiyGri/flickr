package entities

type FlickrImage struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Link     string `json:"link"`
	AuthorID string `json:"author_id"`
	Vote     int    `json:"vote"`
}
