package services

import (
    "errors"
    "log"
    "net/smtp"
    "os"
)

type loginAuth struct {
    username, password string
}

func LoginAuth(username, password string) smtp.Auth {
    return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
    return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
    if more {
        switch string(fromServer) {
        case "Username:":
            return []byte(a.username), nil
        case "Password:":
            return []byte(a.password), nil
        default:
            return nil, errors.New("Unkown fromServer")
        }
    }
    return nil, nil
}

func SendValidationLinkByEmail(toAddress string, hash string) {
    fromAddress := os.Getenv("BOILDERPLATE_EMAIL_ADDRESS")
    auth := LoginAuth(fromAddress, os.Getenv("BOILERPLATE_EMAIL_PASSWORD"))
    to := []string{toAddress}

    // TODO: get the base url from os env
    msg := []byte("To:" + toAddress + "\r\n" +
        "Subject: Access to your Account\r\n" +
        "\r\n" +
        "Click here to access your account: https://app.boilerplate.com?hash=" + hash + "\r\n")

    err := smtp.SendMail("smtp.gmail.com:587", auth, fromAddress, to, msg)

    if err != nil {
        log.Fatal(err)
    }
}
