package data

import "alta/immersive-dashboard-api/features/users"

func ModelToCore(user Users) users.Core{
	return users.Core{
		Id:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Team:      user.Team,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

