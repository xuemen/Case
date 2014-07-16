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
	PatientID int
	Name      string
	Sex       string
	BOD       string
	Address   string
	PMH       string
	FMH       string
	Allergies string
}

const (
	CaseStatusDefault int = 1 << iota
	CaseStatusCreated
	CaseStatusWaitting
	CaseStatusSubmitted
)

func CaseNew(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("path", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	if r.Method == "GET" {
		if len(r.Form["pid"]) > 0 {
			sqlstr := fmt.Sprintf("select PatientID,Name,Sex,ifnull(date(BOD),\"未填写生日\"),ifnull(Address,\"未填写地址\"),ifnull(PMH,\"未填写既往病史\"),ifnull(FMH,\"未填写家族病史\"),ifnull(Allergies,\"未填写过敏史\") from Patient where PatientID = %s", r.Form["pid"][0])
			log.Print("sqlstr:\t", sqlstr)

			db, err := sql.Open("sqlite3", "./case.v0.1.s3db")
			rows, err := db.Query(sqlstr)
			checkErr(err)

			//log.Print(rows)
			var p Patient
			for rows.Next() {
				err = rows.Scan(&p.PatientID, &p.Name, &p.Sex, &p.BOD, &p.Address, &p.PMH, &p.FMH, &p.Allergies)
				checkErr(err)
			}
			log.Print("p:\t", p)
			t, _ := template.ParseFiles("caseedit.gtpl")
			err = t.Execute(w, p)
			log.Print("err:\t", err)
		} else {
			t, _ := template.ParseFiles("caseedit.gtpl")
			t.Execute(w, nil)
		}
	} else if r.Method == "POST" {
		sqlstr := fmt.Sprintf("insert into record (patientid,MainComplaint,ExamReport,Diag,DRR,Presciption,Notes,CreateTime,SubmitTime,IsTemplate,Status) values (%s,\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",datetime(\"now\",\"localtime\"),datetime(\"now\",\"localtime\"),0,%d)",
			r.Form["pid"][0],
			r.Form["MainComplaint"][0],
			r.Form["ExamReport"][0],
			r.Form["Diag"][0],
			r.Form["DRR"][0],
			r.Form["Presciption"][0],
			r.Form["Notes"][0],
			CaseStatusSubmitted)

		log.Print("sqlstr:\t", sqlstr)
		db, err := sql.Open("sqlite3", "./case.v0.1.s3db")
		result, err := db.Exec(sqlstr)
		checkErr(err)
		log.Print("result:\t", result)

		fmt.Fprint(w, "<script>window.location=\"/\"</script>")
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
