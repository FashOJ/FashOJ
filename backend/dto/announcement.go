package dto

type LastAnnouncement struct {
	Abstract string `json:"abstract"`
	Title    string `json:"title"`
}

type CreateAnnouncement struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Announcement struct {
	Id       int    `json:"id"`
	Author   string `json:"author"`
	Avatar   string `json:"avatar"`
	AuthorId int    `json:"author_id"`
	Title    string `json:"title"`
	Abstract string `json:"abstract"`
}

type AnnouncementPage struct {
	Pages         int            `json:"pages"`
	Size          int            `json:"size"`
	Announcements []Announcement `json:"announcements"`
}
