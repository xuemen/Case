package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	CaseStatusDefault int = 1 << iota
	CaseStatusCreated
	CaseStatusWaitting
	CaseStatusSubmitted
)

type RID struct {
	RecordID string
}

type RPID struct {
	RecordID  string
	PatientID string
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
}

func CaseNew(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("path", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	if r.Method == "GET" {
		if len(r.Form["rid"]) > 0 {
			t, _ := template.ParseFiles("static/template/caseedit.gtpl")
			rp := RPID{RecordID: r.Form["rid"][0], PatientID: ""}
			err := t.Execute(w, rp)
			log.Print("err:\t", err)
		} else if len(r.Form["pid"]) > 0 {
			t, _ := template.ParseFiles("static/template/caseedit.gtpl")
			rp := RPID{RecordID: "", PatientID: r.Form["pid"][0]}
			err := t.Execute(w, rp)
			log.Print("err:\t", err)
		} else {
			t, _ := template.ParseFiles("static/template/caseedit.gtpl")
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
		defer db.Close()
		result, err := db.Exec(sqlstr)
		checkErr(err)
		log.Print("result:\t", result)

		fmt.Fprint(w, "<script>window.location=\"/\"</script>")
	}

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
			defer db.Close()
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

			t, _ := template.ParseFiles("static/template/caselist.gtpl")
			err = t.Execute(w, cld)
			log.Print("err:\t", err)
		} else {
			t, _ := template.ParseFiles("static/template/caselist.gtpl")
			t.Execute(w, nil)
		}
	} else if r.Method == "POST" {

	}

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
			t, err := template.ParseFiles("static/template/casedetail.gtpl")
			r := RID{RecordID: r.Form["rid"][0]}
			err = t.Execute(w, r)
			log.Print("err:\t", err)
		} else {
			t, _ := template.ParseFiles("static/template/casedetail.gtpl")
			t.Execute(w, nil)
		}
	} else if r.Method == "POST" {
		sqlstr := fmt.Sprintf("update record set patientid=%s,MainComplaint=\"%s\",ExamReport=\"%s\",Diag=\"%s\",DRR=\"%s\",Presciption=\"%s\",Notes=\"%s\",SubmitTime=datetime(\"now\",\"localtime\"),IsTemplate=0,Status=%d where RecordID=%s",
			r.Form["pid"][0],
			r.Form["MainComplaint"][0],
			r.Form["ExamReport"][0],
			r.Form["Diag"][0],
			r.Form["DRR"][0],
			r.Form["Presciption"][0],
			r.Form["Notes"][0],
			CaseStatusSubmitted,
			r.Form["rid"][0])

		log.Print("sqlstr:\t", sqlstr)
		db, err := sql.Open("sqlite3", "./case.v0.1.s3db")
		defer db.Close()
		result, err := db.Exec(sqlstr)
		checkErr(err)
		log.Print("result:\t", result)

		fmt.Fprint(w, "<script>window.location=\"/\"</script>")
	}

}

func CaseInfo(w http.ResponseWriter, r *http.Request) {
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
			defer db.Close()
			rows, err := db.Query(sqlstr)
			checkErr(err)

			//log.Print(rows)
			var c CaseDetailData

			var hasRes bool
			hasRes = false

			for rows.Next() {
				err = rows.Scan(&c.RecordID, &c.PatientID, &c.MainComplaint, &c.ExamReport, &c.Diag, &c.DRR, &c.Presciption, &c.Notes, &c.CreateTime)
				checkErr(err)
				hasRes = true
			}
			log.Print("c:\t", c)

			if hasRes {
				t, err := template.ParseFiles("static/template/caseinfo.gtpl")
				err = t.Execute(w, c)
				log.Print("err:\t", err)
				err = t.Execute(os.Stdout, c)
			} else {
				fmt.Fprintf(w, "没有这份病历，编号：%s", r.Form["rid"][0])
			}

		} else {
			fmt.Fprint(w, "未指定病历编号。")
		}
	}
}
