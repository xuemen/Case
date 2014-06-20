package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Patient struct {
	PatientID int
	Name      string
	Sex       string
	BOD       string
}

func PatientSearsh(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //解析参数，默认是不会解析的

	if r.Method == "GET" {
		t, _ := template.ParseFiles("patientsearch.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		where := false
		wherestr := ""
		sqlstr := "select PatientID,Name,Sex,BOD from Patient"
		if r.Form["id"][0] != "" {
			if where {
				wherestr = fmt.Sprintf("%s and ID=%s", wherestr, r.Form["id"][0])
			} else {
				wherestr = fmt.Sprintf("%s ID=%s", wherestr, r.Form["id"][0])
			}
			where = true
		}
		if r.Form["name"][0] != "" {
			if where {
				wherestr = fmt.Sprintf("%s and Name=\"%s\"", wherestr, r.Form["name"][0])
			} else {
				wherestr = fmt.Sprintf("%s Name=\"%s\"", wherestr, r.Form["name"][0])
			}

			where = true
		}
		if r.Form["sex"][0] != "" {
			if where {
				wherestr = fmt.Sprintf("%s and Sex=\"%s\"", wherestr, r.Form["sex"][0])
			} else {
				wherestr = fmt.Sprintf("%s Sex=\"%s\"", wherestr, r.Form["sex"][0])
			}

			where = true
		}
		if r.Form["BOD"][0] != "" {
			if where {
				wherestr = fmt.Sprintf("%s and BOD=\"%s\"", wherestr, r.Form["BOD"][0])
			} else {
				wherestr = fmt.Sprintf("%s BOD=\"%s\"", wherestr, r.Form["BOD"][0])
			}

			where = true
		}

		if where {
			sqlstr = fmt.Sprintf("%s where %s limit 10", sqlstr, wherestr)
		} else {
			sqlstr = fmt.Sprintf("%s limit 10", sqlstr)
		}

		db, err := sql.Open("sqlite3", "./case.v0.1.s3db")
		rows, err := db.Query(sqlstr)
		checkErr(err)

		log.Print(rows)

		var parray [10]Patient
		c := parray[0:0]

		for rows.Next() {
			var p Patient

			err = rows.Scan(&p.PatientID, &p.Name, &p.Sex, &p.BOD)
			checkErr(err)

			c = append(c, p)
		}

		t, _ := template.ParseFiles("patientlist.gtpl")
		t.Execute(w, c)
	}

}

func PatientNew(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("Action:", r.Form["action"])
	fmt.Println("ID:", r.Form["id"])
	fmt.Println("Name:", r.Form["name"])
	fmt.Println("Sex:", r.Form["sex"])
	fmt.Println("BOD:", r.Form["BOD"])

	if r.Method == "GET" {
		t, _ := template.ParseFiles("patientnew.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {

	}

}

func PatientUpdate(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("Action:", r.Form["action"])
	fmt.Println("ID:", r.Form["id"])
	fmt.Println("Name:", r.Form["name"])
	fmt.Println("Sex:", r.Form["sex"])
	fmt.Println("BOD:", r.Form["BOD"])

	if r.Method == "GET" {
		t, _ := template.ParseFiles("patientupdate.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {

	}

}
