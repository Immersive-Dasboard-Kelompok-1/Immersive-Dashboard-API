package handler

import "alta/immersive-dashboard-api/features/users"

type RequestUser struct{
	FullName		string		`json:"fullname" form:"fullname"`			
	Email			string		`json:"email" form:"email"`				
	Password		string		`json:"password" form:"password"`				
	Team 			string 		`json:"team" form:"team"`				
	Role			string 		`json:"role" form:"role"`		
	Status			string 		`json:"status" form:"status"`				
}

func RequestToCoreUser(input RequestUser)users.Core{
	return users.Core{
		FullName: input.FullName,
		Email: input.Email,
		Password: input.Password,
		Team: input.Team,
		Role: input.Role,
		Status: input.Status,
	}
}

