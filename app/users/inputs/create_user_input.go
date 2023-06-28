package inputs


type CreateUserInput struct {
	Name string `json:"name"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	PhonePrefix string `json:"phonePrefix" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	CountryCode string `json:"countryCode" validate:"required"`
}
