package models

type Comment struct {
	Id 					int		`json:"id"`
	Content 			string	`json:"content"`
	Created_by_user_id 	int		`json:"created_by_user_id"`
	Ticket_id 			int		`json:"ticket_id"`
	Creation_date 		string	`json:"creation_date"`
	Last_update_date 	string	`json:"last_update_date"`
}