package models

import "app"

type Picture struct {
	Id        int64    `json:"id"`
	Title     string   `json:"title"`
	CateId    int64    `json:"cate_id"`
	AuthorId  int64    `json:"author_id"`
	Sort      int64    `json:"sort"`
	IsComment bool     `json:"is_comment"`
	Img       string   `json:"img"`
	Source    string   `json:"source"`
	Keywords  string   `json:"keywords"`
	Intro     string   `json:"intro"`
	Images    string   `json:"images"`
	StartTime string   `json:"start_time"`
	EndTime   string   `json:"end_time"`
	Status    int64    `json:"status"`
	CreatedAt app.Time `json:"created_at"`
	UpdatedAt app.Time `json:"updated_at"`
}
