package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Patient struct {
	PatientID  int
	Name       string
	Sex        string
	BOD        string
	Address    string
	CreateTime string
	Diag       string
}

func PatientSearsh(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("path", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	if r.Method == "GET" {
		t, _ := template.ParseFiles("patientsearch.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		f := r.Form
		log.Print("r.Form:\t", r.Form)
		log.Print("r.Form[\"b\"]:\t", r.Form["b"])
		log.Print("r.Form[\"b\"][0]:\t", r.Form["b"][0])

		where := false
		wherestr := ""
		sqlstr := "select Patient.*,datetime(CreateTime),ifnull(Diag,'未填写诊断') from Patient left join Record on Patient.PatientID = Record.PatientID %s group by Patient.PatientID limit 10"
		if r.Form["id"][0] != "" {
			if where {
				wherestr = fmt.Sprintf("%s and Patient.PatientID=%s", wherestr, r.Form["id"][0])
			} else {
				wherestr = fmt.Sprintf("%s Patient.PatientID=%s", wherestr, r.Form["id"][0])
			}
			where = true
			log.Print("wherestr:\t", wherestr)
		}

		if r.Form["name"][0] != "" {
			if where {
				wherestr = fmt.Sprintf("%s and Patient.Name=\"%s\"", wherestr, r.Form["name"][0])
			} else {
				wherestr = fmt.Sprintf("%s Patient.Name=\"%s\"", wherestr, r.Form["name"][0])
			}

			where = true
			log.Print("wherestr:\t", wherestr)
		}

		if r.Form["sex"][0] != "All" {
			if where {
				wherestr = fmt.Sprintf("%s and Patient.Sex=\"%s\"", wherestr, r.Form["sex"][0])
			} else {
				wherestr = fmt.Sprintf("%s Patient.Sex=\"%s\"", wherestr, r.Form["sex"][0])
			}

			where = true
			log.Print("wherestr:\t", wherestr)
		}

		if r.Form["BOD"][0] != "" {
			if where {
				wherestr = fmt.Sprintf("%s and Patient.BOD=\"%s\"", wherestr, r.Form["BOD"][0])
			} else {
				wherestr = fmt.Sprintf("%s Patient.BOD=\"%s\"", wherestr, r.Form["BOD"][0])
			}

			where = true
			log.Print("wherestr:\t", wherestr)
		}
		log.Print("f.Get(b):\t", f.Get("b"))
		if f.Get("b") == "24小时内就诊" {
			if where {
				wherestr = fmt.Sprintf("%s and (strftime('%%s','now') - strftime('%%s',createtime))<86400", wherestr)
			} else {
				wherestr = fmt.Sprintf("%s (strftime('%%s','now') - strftime('%%s',createtime))<86400", wherestr)
			}

			where = true
			log.Print("wherestr:\t", wherestr)
		}

		if f.Get("b") == "7天内就诊" {
			if where {
				wherestr = fmt.Sprintf("%s and (strftime('%%s','now') - strftime('%%s',createtime))<604800", wherestr)
			} else {
				wherestr = fmt.Sprintf("%s (strftime('%%s','now') - strftime('%%s',createtime))<604800", wherestr)
			}

			where = true
			log.Print("wherestr:\t", wherestr)
		}

		if where {
			wherestr = fmt.Sprintf("where %s", wherestr)
		}

		log.Print("sqlstr:\t", sqlstr)
		sqlstr = fmt.Sprintf(sqlstr, wherestr)
		log.Print("sqlstr:\t", sqlstr)

		db, err := sql.Open("sqlite3", "./case.v0.1.s3db")
		rows, err := db.Query(sqlstr)
		checkErr(err)

		log.Print(rows)

		var parray [10]Patient
		c := parray[0:0]

		HasResult := false

		for rows.Next() {
			var p Patient

			err = rows.Scan(&p.PatientID, &p.Name, &p.Sex, &p.BOD, &p.Address, &p.CreateTime, &p.Diag)
			checkErr(err)

			c = append(c, p)
			HasResult = true
			log.Print("c:\t", c)
			log.Print("p:\t", p)
		}

		t, _ := template.ParseFiles("patientsearch.gtpl")
		if HasResult {
			t.Execute(w, c)
		} else {
			t.Execute(w, nil)
		}

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
