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

type Case struct {
	RecordID      int
	MainComplaint string
	ExamReport    string
	Diag          string
	DRR           string
	Presciption   string
	CreateTime    string
}

type CaseListData struct {
	PatientID int
	Name      string
	Sex       string
	BOD       string
	Address   string
	Cases     []Case
}

func CaseList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("path", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	if r.Method == "GET" {
		if len(r.Form["pid"]) > 0 {
			sqlstr := fmt.Sprintf("select PatientID,Name,Sex,ifnull(date(BOD),\"未填写生日\"),ifnull(Address,\"未填写地址\") from Patient where PatientID = %s", r.Form["pid"][0])
			log.Print("sqlstr:\t", sqlstr)

			db, err := sql.Open("sqlite3", "./case.v0.1.s3db")
			rows, err := db.Query(sqlstr)
			checkErr(err)

			//log.Print(rows)
			var cld CaseListData
			for rows.Next() {
				err = rows.Scan(&cld.PatientID, &cld.Name, &cld.Sex, &cld.BOD, &cld.Address)
				checkErr(err)
			}
			log.Print("cld:\t", cld)

			sqlstr = fmt.Sprintf("Select RecordID,ifnull(MainComplaint,\"未填写主诉\"),ifnull(ExamReport,\"无检验报告\"),ifnull(Diag,\"未填写诊断\"),ifnull(DRR,\"未填写医嘱\"),ifnull(Presciption,\"未填写处方\"),ifnull(datetime(CreateTime),\"未提交\") from Record where PatientID = %s limit 100", r.Form["pid"][0])
			log.Print("sqlstr:\t", sqlstr)

			rows, err = db.Query(sqlstr)
			checkErr(err)

			//log.Print(rows)
			var CArray [100]Case
			cld.Cases = CArray[0:0]

			for rows.Next() {
				var c Case

				err = rows.Scan(&c.RecordID, &c.MainComplaint, &c.ExamReport, &c.Diag, &c.DRR, &c.Presciption, &c.CreateTime)
				checkErr(err)

				cld.Cases = append(cld.Cases, c)
			}

			t, _ := template.ParseFiles("caselist.gtpl")
			err = t.Execute(w, cld)
			log.Print("err:\t", err)
		} else {
			t, _ := template.ParseFiles("caselist.gtpl")
			t.Execute(w, nil)
		}
	} else if r.Method == "POST" {

	}

}

type CaseDetailData struct {
	RecordID      int
	PatientID     int
	MainComplaint string
	ExamReport    string
	Diag          string
	DRR           string
	Presciption   string
	Notes         string
	CreateTime    string
	Name          string
	Sex           string
	BOD           string
	Address       string
	PMH           string
	FMH           string
	Allergies     string
}

func CaseDetail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("path", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	if r.Method == "GET" {
		if len(r.Form["rid"]) > 0 {
			sqlstr := fmt.Sprintf("Select RecordID,PatientID,ifnull(MainComplaint,\"未填写主诉\"),ifnull(ExamReport,\"无检验报告\"),ifnull(Diag,\"未填写诊断\"),ifnull(DRR,\"未填写医嘱\"),ifnull(Presciption,\"未填写处方\"),ifnull(Notes,\"未填写备注\"),ifnull(datetime(CreateTime),\"未提交\") from Record where RecordID = %s", r.Form["rid"][0])
			log.Print("sqlstr:\t", sqlstr)

			db, err := sql.Open("sqlite3", "./case.v0.1.s3db")
			rows, err := db.Query(sqlstr)
			checkErr(err)

			//log.Print(rows)
			var c CaseDetailData
			for rows.Next() {
				err = rows.Scan(&c.RecordID, &c.PatientID, &c.MainComplaint, &c.ExamReport, &c.Diag, &c.DRR, &c.Presciption, &c.Notes, &c.CreateTime)
				checkErr(err)
			}
			log.Print("c:\t", c)

			sqlstr = fmt.Sprintf("select PatientID,Name,Sex,ifnull(date(BOD),\"未填写生日\"),ifnull(Address,\"未填写地址\"),ifnull(PMH,\"未填写既往病史\"),ifnull(FMH,\"未填写家族病史\"),ifnull(Allergies,\"未填写过敏史\") from Patient where PatientID = %d", c.PatientID)
			log.Print("sqlstr:\t", sqlstr)

			rows, err = db.Query(sqlstr)
			checkErr(err)
			log.Print("rows:\t", rows)

			for rows.Next() {
				err = rows.Scan(&c.PatientID, &c.Name, &c.Sex, &c.BOD, &c.Address, &c.PMH, &c.FMH, &c.Allergies)
				checkErr(err)
			}
			log.Print("c:\t", c)
			t, err := template.ParseFiles("casedetail.gtpl")
			checkErr(err)
			err = t.Execute(w, c)
			checkErr(err)
			log.Print("err:\t", err)
		} else {
			t, _ := template.ParseFiles("casedetail.gtpl")
			t.Execute(w, nil)
		}
	} else if r.Method == "POST" {

	}

}
