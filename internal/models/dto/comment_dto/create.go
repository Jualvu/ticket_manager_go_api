package comment_dto

type CreateCommentRequest struct {
	Content 			string	`json:"content"`
	Created_by_user_id 	int		`json:"created_by_user_id"`
	Ticket_id 			int		`json:"ticket_id"`
}