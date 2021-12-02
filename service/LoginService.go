package service

type LoginService interface {
	LoginUser(email string, password string) bool
}

type loginInformation struct {
	email      string
	password   string
}

func