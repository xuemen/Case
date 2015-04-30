package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func RecptDetail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("path", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	if r.Method == "GET" {
		t, _ := template.ParseFiles("static/template/recptdetail.html")
		if len(r.Form["rtid"]) > 0 {
			rtid, _ := strconv.Atoi(r.Form["rtid"][0])
			rt := rtslice[rtid]
			log.Print(rt)
			t.Execute(w, rt)
		}
	}
}
