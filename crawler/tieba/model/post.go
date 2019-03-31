package model

type Floor struct {
	Content  string
	Author   string
	AuthorID string
}

type Post struct {
	PageIndex int
	Url       string
	Title     string
	Floors    []Floor
}

func (p *Post) String() string {
	return ""
}
