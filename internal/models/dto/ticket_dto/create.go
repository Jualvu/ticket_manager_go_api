package ticket_dto

type CreateTicketRequest struct {
	Title 					string	`json:"title"` 
	Description 			string	`json:"description"`
	State_id 				int		`json:"state_id"`
	Priority_id 			int		`json:"priority_id"`
	Assigned_to_user_id 	int		`json:"assigned_to_user_id"`
	Created_by_user_id 		int		`json:"created_by_user_id"`
}