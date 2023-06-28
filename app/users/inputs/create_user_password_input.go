package inputs

import (
	"time"
)


type CreateUserPasswordInput struct {
	UserId uint64
	Hash string
	ExpiresAt time.Time
}
