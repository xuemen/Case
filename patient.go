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
		where := false
		wherestr := ""
		sqlstr := "select Patient.*,ifnull(datetime(CreateTime),'未就诊'),ifnull(Diag,'未填写诊断') from Patient left join Record on Patient.PatientID = Record.PatientID %s group by Patient.PatientID"
		if len(r.Form["id"]) > 0 {
			if where {
				wherestr = fmt.Sprintf("%s and Patient.PatientID=%s", wherestr, r.Form["id"][0])
			} else {
				wherestr = fmt.Sprintf("%s Patient.PatientID=%s", wherestr, r.Form["id"][0])
			}
			where = true
			log.Print("wherestr:\t", wherestr)
		}

		if len(r.Form["name"]) > 0 {
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

		if len(r.Form["BOD"]) > 0 {
			if where {
				wherestr = fmt.Sprintf("%s and Patient.BOD=\"%s\"", wherestr, r.Form["BOD"][0])
			} else {
				wherestr = fmt.Sprintf("%s Patient.BOD=\"%s\"", wherestr, r.Form["BOD"][0])
			}

			where = true
			log.Print("wherestr:\t", wherestr)
		}

		if r.Form["time"][0] == "24h" {
			if where {
				wherestr = fmt.Sprintf("%s and (strftime('%%s','now') - strftime('%%s',createtime))<86400", wherestr)
			} else {
				wherestr = fmt.Sprintf("%s (strftime('%%s','now') - strftime('%%s',createtime))<86400", wherestr)
			}

			where = true
			log.Print("wherestr:\t", wherestr)
		}

		if r.Form["time"][0] == "7d" {
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
	fmt.Println("path", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	if r.Method == "GET" {
		t, _ := template.ParseFiles("patientnew.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		if (len(r.Form["id"][0]) == 0) || (len(r.Form["name"][0]) == 0) {
			fmt.Fprint(w, "<script>alert(\"必须填写编号和姓名\");</script>")
			t, _ := template.ParseFiles("patientnew.gtpl")
			t.Execute(w, nil)
		}
		sqlstr := fmt.Sprintf("insert into patient (patientid,name,sex,bod,address,PMH,FMH,Allergies,FVT) values (\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",datetime(\"now\",\"localtime\"))",
			r.Form["id"][0],
			r.Form["name"][0],
			r.Form["sex"][0],
			r.Form["BOD"][0],
			r.Form["Address"][0],
			r.Form["PMH"][0],
			r.Form["FMH"][0],
			r.Form["Allergies"][0])

		log.Print("sqlstr:\t", sqlstr)
		db, err := sql.Open("sqlite3", "./case.v0.1.s3db")
		result, err := db.Exec(sqlstr)
		checkErr(err)
		log.Print("result:\t", result)

		fmt.Fprint(w, "<script>window.open(\"/case/new\")</script>")
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

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	if r.Method == "GET" {
		t, _ := template.ParseFiles("patientupdate.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {

	}

}
