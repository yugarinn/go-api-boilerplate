package core

import (
	"github.com/yugarinn/go-api-boilerplate/lib"
)


type App struct {
	TwilioClient lib.TwilioClientInterface
}

func BootstrapApplication() *App {
    app := App{
        TwilioClient: &lib.TwilioClient{},
    }

	return &app
}
