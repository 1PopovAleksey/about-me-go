package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name       string
	LastName   string
	Age        uint16
	Stack      string
	Experience string
	Hobbies    []string
	Telegram   string
	Gmail      string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("My name is: %s %s.\n"+
		"I'm: %d.\n"+
		"My full stack: %s.\n"+
		"My experience in years: %s",
		u.Name, u.LastName, u.Age, u.Stack, u.Experience)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, "")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	aleksey := User{
		Name:       "Alex",
		LastName:   "Popov",
		Age:        18,
		Stack:      "JavaScript, TypeScript, React, GO, Python",
		Experience: "0.2+",
		Hobbies:    []string{"VideoGame", "Amine", "Photo", "Music"},
	}
	aleksey.setNewName("Aleksey")
	//fmt.Fprintf(w, aleksey.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/about-me.html")
	tmpl.Execute(w, aleksey)
}

func contactsPages(w http.ResponseWriter, r *http.Request) {
	contacts := User{
		Telegram: "@AlekseyPopovDev",
		Gmail:    "AlekseyPopov.Dev@gmail.com",
	}
	tmpl, _ := template.ParseFiles("templates/contacts.html")
	tmpl.Execute(w, contacts)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about/", aboutPage)
	http.HandleFunc("/contacts/", contactsPages)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
