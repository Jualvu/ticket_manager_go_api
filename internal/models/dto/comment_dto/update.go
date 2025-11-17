package comment_dto

type UpdateCommentRequest struct {
	Id 					int		`json:"id"`
	Content 			string	`json:"content"`
}