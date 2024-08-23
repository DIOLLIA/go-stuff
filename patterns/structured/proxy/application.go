package main

import "net/http"

type Application struct {
}

func (a *Application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == http.MethodGet {
		return 201, "OK"
	}
	if url == "/create/user" && method == http.MethodPost {
		return 200, "OK"
	}
	return 404, "NotOK"
}
