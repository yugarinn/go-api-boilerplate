package responses

import (
	models "github.com/yugarinn/go-api-boilerplate/app/users/models"
)

type UserResponse struct {
	ID uint64          `json:"id"`
	Email string       `json:"email"`
	Name string        `json:"name"`
	LastName string    `json:"lastName"`
	CountryCode string `json:"countryCode"`
	PhoneNumber string `json:"phoneNumber"`
	PhonePrefix string `json:"phonePrefix"`
}

type PaginatedUserResponse struct {
	Users []UserResponse  `json:"users"`
	Pagination Pagination `json:"pagination"`
}

func SerializeUser(user models.User) UserResponse {
	return UserResponse{
		ID: user.ID,
		Email: user.Email,
		Name: user.Name,
		LastName: user.LastName,
		CountryCode: user.CountryCode,
		PhoneNumber: user.PhoneNumber,
		PhonePrefix: user.PhonePrefix,
	}
}
