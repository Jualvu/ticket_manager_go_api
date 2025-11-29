package models

type Ticket struct {
	Id 						int		
	Title 					string	 
	Description 			string	
	State_id 				int		
	Priority_id 			int		
	Assigned_to_user_id 	int		
	Created_by_user_id 		int		
	Creation_date 			string	
	Last_update_date 		string	
}