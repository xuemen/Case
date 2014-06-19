package main

import (
	"fmt"
	"html/template"

	"net/http"
)

func CaseNew(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("Action:", r.Form["action"])
	fmt.Println("ID:", r.Form["id"])
	fmt.Println("Name:", r.Form["name"])
	fmt.Println("Sex:", r.Form["sex"])
	fmt.Println("BOD:", r.Form["BOD"])

	if r.Method == "GET" {
		t, _ := template.ParseFiles("casenew.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {

	}

}

func CaseList(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("Action:", r.Form["action"])
	fmt.Println("ID:", r.Form["id"])
	fmt.Println("Name:", r.Form["name"])
	fmt.Println("Sex:", r.Form["sex"])
	fmt.Println("BOD:", r.Form["BOD"])

	if r.Method == "GET" {
		t, _ := template.ParseFiles("caselist.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {

	}

}

func CaseDetail(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("Action:", r.Form["action"])
	fmt.Println("ID:", r.Form["id"])
	fmt.Println("Name:", r.Form["name"])
	fmt.Println("Sex:", r.Form["sex"])
	fmt.Println("BOD:", r.Form["BOD"])

	if r.Method == "GET" {
		t, _ := template.ParseFiles("casedetial.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {

	}

}
