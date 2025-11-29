package models

type Comment struct {
	Id 					int		
	Content 			string	
	Created_by_user_id 	int		
	Ticket_id 			int		
	Creation_date 		string	
	Last_update_date 	string	
}