package main

import (
	"fmt"
	"math/rand"
	"time"
)

type LPair struct {
	login, pwd string
}

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-")

func init() {
	rand.NewSource(time.Now().UnixNano())
}
func createPair() (login, pwd string) {
	var loginLength = 8
	domains := [5]string{"gmail.com", "yandex.ru", "mail.ru", "yahoo.de", "hotmail.fi"}

	loginTemplate := make([]rune, loginLength)
	for i := range loginTemplate {
		loginTemplate[i] = runes[rand.Intn(len(runes))]
	}
	login = string(loginTemplate)
	time.Sleep(1 * time.Millisecond)
	log, Pwd := fmt.Sprintf("login: %s@%s", login, domains[rand.Intn(4)]), "password"
	return log, Pwd
}
