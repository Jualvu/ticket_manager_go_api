package ticket_dto

type UpdateTicketRequest struct {
	Id 						int		`json:"id"`
	Title 					string	`json:"title"` 
	Description 			string	`json:"description"`
	State_id 				int		`json:"state_id"`
	Priority_id 			int		`json:"priority_id"`
	Assigned_to_user_id 	int		`json:"assigned_to_user_id"`
}