package data

import "alta/immersive-dashboard-api/features/users"

func CoreToUser(data users.Core, ) Users{
	return Users{
		FullName: data.FullName,
		Email:    data.Email,
		Password: data.Password,
		Team:     data.Team,
		Role:     data.Role,
		Status:   data.Status,	
	}
}