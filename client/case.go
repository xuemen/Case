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
	Cases []Case
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
	ReadOnly      string
}

type FourDiag struct {
	CaseID int
	StrA1  string //问诊-主诉
	StrA2  string //问诊-现病史
	StrA3  string //问诊-既往史
	StrA4  string //问诊-过敏史
	StrA5  string //问诊-婚育史
	StrA6  string //问诊-个人史
	StrA7  string //问诊-家族史

	StrB1 string //望诊-舌诊
	StrB2 string //望诊-神色形态
	StrB3 string //望诊-舌诊图片
	StrB4 string //望诊-胸腹
	StrB5 string //望诊-腰背四肢爪甲
	StrB6 string //望诊-皮肤毛发
	StrB7 string //望诊-头面五官颈项
	StrB8 string //望诊-前后二阴及排泄物

	StrC1 string //闻切诊-闻诊
	StrC2 string //闻切诊-脉诊
	StrC3 string //闻切诊-其它

	StrD1 string //摘要-症状
	StrD2 string //摘要-舌诊
	StrD3 string //摘要-脉诊
}

type Exam struct {
	CaseID int
	StrA1  string //体格检查
	StrB1  string //实验室检查-血常规
	StrB2  string //实验室检查-心电图
	StrB3  string //实验室检查-尿常规
	StrB4  string //实验室检查-CT
	StrB5  string //实验室检查-大便常规
	StrB6  string //实验室检查-MRI
	StrB7  string //实验室检查-血生化
	StrB8  string //实验室检查-超声
	StrB9  string //实验室检查-X光片
	StrB10 string //实验室检查-其它
}

type DiagAndTreatment struct {
	CaseID int

	StrA1 string //诊断与治法-中医疾病
	StrA2 string //诊断与治法-西医疾病
	StrA3 string //诊断与治法-中医证候
	StrA4 string //诊断与治法-治则执法

	Barray []Recpt //中医处方

	StrC1 string //穴位处方

	StrD1 string //其它治疗-中成药
	StrD2 string //其它治疗-西成药
	StrD3 string //其它治疗-其它

}

type Recpt struct {
	RecptID   int
	Name      string
	RType     string
	Amount    int
	CreatorID int
	Detail    []Drug
}

type Drug struct {
	Name       string
	Amount     float64
	Role       string
	Processing string
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
		var cld CaseListData
		var sqlstr string
		if len(r.Form["pid"]) > 0 {
			sqlstr = fmt.Sprintf("Select RecordID,ifnull(MainComplaint,\"未填写主诉\"),ifnull(ExamReport,\"无检验报告\"),ifnull(Diag,\"未填写诊断\"),ifnull(DRR,\"未填写医嘱\"),ifnull(Presciption,\"未填写处方\"),ifnull(datetime(CreateTime),\"未提交\") from Record where PatientID = %s limit 100", r.Form["pid"][0])
		} else {
			sqlstr = "Select RecordID,ifnull(MainComplaint,\"未填写主诉\"),ifnull(ExamReport,\"无检验报告\"),ifnull(Diag,\"未填写诊断\"),ifnull(DRR,\"未填写医嘱\"),ifnull(Presciption,\"未填写处方\"),ifnull(datetime(CreateTime),\"未提交\") from Record limit 100"
		}
		log.Print("sqlstr:\t", sqlstr)

		db, err := sql.Open("sqlite3", "./case.v0.1.s3db")
		defer db.Close()
		rows, err := db.Query(sqlstr)
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
			if len(r.Form["readonly"]) > 0 {
				c.ReadOnly = "true"
			}

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
