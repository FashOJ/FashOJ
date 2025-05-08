package dto

type LastAnnouncement struct {
	Abstract string `json:"abstract"`
	Title    string `json:"title"`
}

type CreateAnnouncement struct {
	Title string `json:"title"`
	Content string `json:"content"`
}