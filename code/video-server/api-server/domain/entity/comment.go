package entity

type Comment struct {
	Id        string
	Video_id  string
	Author_id int
	Content   string
}

type CommentDTO struct {
	Id      string
	VideoId string
	Author  string
	Content string
}
