package factories

import (
	"gorm.io/gorm"

	"github.com/yugarinn/go-api-boilerplate/connections"
)

var database *gorm.DB = connections.Database()
