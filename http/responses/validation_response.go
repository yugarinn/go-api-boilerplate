package responses

import (
	"time"

	models "github.com/yugarinn/go-api-boilerplate/app/users/models"
)

type UserValidationResponse struct {
	ID uint64           `json:"id"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func SerializeValidation(validation models.UserValidation) UserValidationResponse {
	return UserValidationResponse{
		ID: validation.ID,
		ExpiresAt: validation.ExpiresAt,
	}
}
