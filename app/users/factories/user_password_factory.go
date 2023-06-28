package factories

import (
	"time"

	"github.com/yugarinn/go-api-boilerplate/app/users/inputs"
	"github.com/yugarinn/go-api-boilerplate/app/users/models"
	"github.com/jaswdr/faker"
)

func CreateUserPassword(properties inputs.CreateUserPasswordInput) users.UserPassword {
    password := users.UserPassword{UserId: properties.UserId, Hash: properties.Hash, ExpiresAt: properties.ExpiresAt}

	if password.UserId == 0 {
		user := CreateUser(UserFactoryInput{})
		password.UserId = user.ID
	}

	if len(password.Hash) == 0 {
		password.Hash = faker.New().RandomStringWithLength(16)
	}

	if password.ExpiresAt.IsZero() {
		password.ExpiresAt, _ = time.Parse(time.RFC3339, "2099-01-01T00:00:00Z")
	}

	database.Create(&password)

	return password
}
