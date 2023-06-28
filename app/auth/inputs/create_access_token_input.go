package inputs

type CreateAccessTokenInput struct {
	PasswordHash string `json:"passwordHash" validate:"required"`
}
