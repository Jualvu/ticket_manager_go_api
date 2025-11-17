package models

type Ticket struct {
	Id 						int		`json:"id"`
	Title 					string	`json:"title"` 
	Description 			string	`json:"description"`
	State_id 				int		`json:"state_id"`
	Priority_id 			int		`json:"priority_id"`
	Assigned_to_user_id 	int		`json:"assigned_to_user_id"`
	Created_by_user_id 		int		`json:"created_by_user_id"`
	Creation_date 			string	`json:"creation_date"`
	Last_update_date 		string	`json:"last_update_date"`
}