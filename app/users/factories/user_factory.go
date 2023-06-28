package factories

import (
	"github.com/yugarinn/go-api-boilerplate/app/users/models"
	"github.com/jaswdr/faker"
)


type UserFactoryInput struct {
	Name string
	LastName string
	Email string
	PhonePrefix string
	PhoneNumber string
	CountryCode string
	IsConfirmed bool
}

func CreateUser(properties UserFactoryInput) users.User {
	faker := faker.New()
    user := users.User{
		Email: properties.Email,
		Name: properties.Name,
		LastName: properties.LastName,
		PhoneNumber: properties.PhoneNumber,
		CountryCode: properties.CountryCode,
		IsConfirmed: properties.IsConfirmed,
	}

	if len(user.Email) == 0 {
		user.Email = faker.Person().Contact().Email
	}

	if len(user.PhoneNumber) == 0 {
		user.PhoneNumber = faker.Person().Contact().Phone
	}

	if len(user.CountryCode) == 0 {
		user.CountryCode = "ES"
	}

	if len(user.Name) == 0 {
		user.Name = faker.Person().Name()
	}

	if len(user.LastName) == 0 {
		user.Name = faker.Person().LastName()
	}

	database.Create(&user)

	return user
}

func CreateUsersList(number int, properties UserFactoryInput) []users.User {
	var users []users.User

	for i := 0; i < number; i++ {
		users = append(users, CreateUser(properties))
	}

	return users
}
