package main

import (
	"fmt"
	"net/http"
)

type Repository struct {
	username string
	password string
}

func handleRequest (w http.ResponseWriter, r *http.Request){
	repo := []Repository{
		{username: "user1", password: "123456"},
		{username: "user2", password: "456789"},
		{username: "user3", password: "789123"},
	}

	u, p, ok := r.BasicAuth()
	if !ok {
		fmt.Println("Ошибка BasicAuth")
		w.WriteHeader(401)
		return
	}
	for _, users := range repo {
		if u == users.username && p == users.password {
			fmt.Println("Авторизация успешна. ")
			w.WriteHeader(200)
			return
		}
	}
	fmt.Println("Неверный логин или пароль.")
	w.WriteHeader(401)
	return
}

func main (){
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/authserver", handler)
	http.ListenAndServe(":8060", nil)
}