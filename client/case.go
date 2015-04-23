package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

type Record struct {
	RecordID             int
	PatientID            int
	FourDiagInfo         FourDiag
	ExamInfo             Exam
	DiagAndTreatmentInfo DiagAndTreatment
	CreateTime           string
	ReadOnly             string
}

type FourDiag struct {
	RecordID int
	StrA1    string //问诊-主诉
	StrA2    string //问诊-现病史
	StrA3    string //问诊-既往史
	StrA4    string //问诊-过敏史
	StrA5    string //问诊-婚育史
	StrA6    string //问诊-个人史
	StrA7    string //问诊-家族史

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
	RecordID int
	StrA1    string //体格检查
	StrB1    string //实验室检查-血常规
	StrB2    string //实验室检查-心电图
	StrB3    string //实验室检查-尿常规
	StrB4    string //实验室检查-CT
	StrB5    string //实验室检查-大便常规
	StrB6    string //实验室检查-MRI
	StrB7    string //实验室检查-血生化
	StrB8    string //实验室检查-超声
	StrB9    string //实验室检查-X光片
	StrB10   string //实验室检查-其它
}

type DiagAndTreatment struct {
	RecordID int

	StrA1 string //诊断与治法-中医疾病
	StrA2 string //诊断与治法-西医疾病
	StrA3 string //诊断与治法-中医证候
	StrA4 string //诊断与治法-治则执法

	RtArray []Recpt //中医处方

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

func DatEdit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("path", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	if r.Method == "GET" {
		t, _ := template.ParseFiles("static/template/datedit.html")
		if len(r.Form["rid"]) > 0 {
			rid, _ := strconv.Atoi(r.Form["rid"][0])
			rc := rcslice[rid]
			log.Print(rc)
			t.Execute(w, rc.DiagAndTreatmentInfo)
		}
	} else if r.Method == "POST" {
		var rid int
		var dat DiagAndTreatment
		var rc Record
		if len(r.Form["rid"]) > 0 {
			rid, _ = strconv.Atoi(r.Form["rid"][0])

			if len(r.Form["StrA1"]) > 0 {
				dat.StrA1 = r.Form["StrA1"][0]
				dat.StrA2 = r.Form["StrA2"][0]
				dat.StrA3 = r.Form["StrA3"][0]
				dat.StrA4 = r.Form["StrA4"][0]

				dat.RtArray = []Recpt{rtslice[1], rtslice[1], rtslice[1]}
				dat.StrC1 = r.Form["StrC1"][0]

				dat.StrD1 = r.Form["StrD1"][0]
				dat.StrD2 = r.Form["StrD2"][0]
				dat.StrD3 = r.Form["StrD3"][0]
			}
			rc = rcslice[rid]
			dat.RecordID = rid
			rc.DiagAndTreatmentInfo = dat
			saverecord(rid, rc)
			fmt.Fprint(w, "<script>window.location=\"/\"</script>")
		}
	}
}

func ExamEdit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("path", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	if r.Method == "GET" {
		t, _ := template.ParseFiles("static/template/examedit.html")
		if len(r.Form["rid"]) > 0 {
			rid, _ := strconv.Atoi(r.Form["rid"][0])
			rc := rcslice[rid]
			log.Print(rc)
			t.Execute(w, rc.ExamInfo)
		}
	} else if r.Method == "POST" {
		var rid int
		var ex Exam
		var rc Record
		if len(r.Form["rid"]) > 0 {
			rid, _ = strconv.Atoi(r.Form["rid"][0])

			if len(r.Form["StrA1"]) > 0 {
				ex.StrA1 = r.Form["StrA1"][0]

				ex.StrB1 = r.Form["StrB1"][0]
				ex.StrB2 = r.Form["StrB2"][0]
				ex.StrB3 = r.Form["StrB3"][0]
				ex.StrB4 = r.Form["StrB4"][0]
				ex.StrB5 = r.Form["StrB5"][0]
				ex.StrB6 = r.Form["StrB6"][0]
				ex.StrB7 = r.Form["StrB7"][0]
				ex.StrB8 = r.Form["StrB8"][0]
				ex.StrB9 = r.Form["StrB9"][0]
				ex.StrB10 = r.Form["StrB10"][0]
			}
			rc = rcslice[rid]
			ex.RecordID = rid
			rc.ExamInfo = ex
			saverecord(rid, rc)
			fmt.Fprintf(w, "<script>window.location=\"/case/exam?rid=%d\"</script>", rid)
		}
	}
}

func FourDiagEdit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("path", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	cur := time.Now()
	curstr := cur.Format("2006-01-02 15:04:05")

	if r.Method == "GET" {
		t, _ := template.ParseFiles("static/template/fourdiagedit.html")
		if len(r.Form["rid"]) > 0 {
			rid, _ := strconv.Atoi(r.Form["rid"][0])
			rc := rcslice[rid]
			log.Print(rc)
			t.Execute(w, rc.FourDiagInfo)
		} else if len(r.Form["pid"]) > 0 {
			//pid, _ := strconv.Atoi(r.Form["pid"][0])
			fd := FourDiag{0, "问诊-主诉", "问诊-现病史", "问诊-既往史", "问诊-过敏史", "问诊-婚育史", "问诊-个人史", "问诊-家族史", "望诊-舌诊", "望诊-神色形态", "望诊-舌诊图片", "望诊-胸腹", "望诊-腰背四肢爪甲", "望诊-皮肤毛发", "望诊-头面五官颈项", "望诊-前后二阴及排泄物", "闻切诊-闻诊", "闻切诊-脉诊", "闻切诊-其它", "摘要-症状", "摘要-舌诊", "摘要-脉诊"}
			t.Execute(w, fd)
		} else {
			t.Execute(w, nil)
		}
	} else if r.Method == "POST" {
		var rid, pid int
		var fd FourDiag
		var rc Record
		if len(r.Form["rid"]) > 0 {
			rid, _ = strconv.Atoi(r.Form["rid"][0])
		} else {
			rid = getnewrid()
		}
		fd.RecordID = rid
		rc.RecordID = rid

		if len(r.Form["pid"]) > 0 {
			pid, _ = strconv.Atoi(r.Form["pid"][0])

			rc.PatientID = pid
			if len(r.Form["StrA1"]) > 0 {
				fd.StrA1 = r.Form["StrA1"][0]
				fd.StrA2 = r.Form["StrA2"][0]
				fd.StrA3 = r.Form["StrA3"][0]
				fd.StrA4 = r.Form["StrA4"][0]
				fd.StrA5 = r.Form["StrA5"][0]
				fd.StrA6 = r.Form["StrA6"][0]
				fd.StrA7 = r.Form["StrA7"][0]

				fd.StrB1 = r.Form["StrB1"][0]
				fd.StrB2 = r.Form["StrB2"][0]
				fd.StrB3 = r.Form["StrB3"][0]
				fd.StrB4 = r.Form["StrB4"][0]
				fd.StrB5 = r.Form["StrB5"][0]
				fd.StrB6 = r.Form["StrB6"][0]
				fd.StrB7 = r.Form["StrB7"][0]
				fd.StrB8 = r.Form["StrB8"][0]

				fd.StrC1 = r.Form["StrC1"][0]
				fd.StrC2 = r.Form["StrC2"][0]
				fd.StrC3 = r.Form["StrC3"][0]

				fd.StrD1 = r.Form["StrD1"][0]
				fd.StrD2 = r.Form["StrD2"][0]
				fd.StrD3 = r.Form["StrD3"][0]
			}

			rc.FourDiagInfo = fd
			rc.CreateTime = curstr

			saverecord(rid, rc)
			fmt.Fprintf(w, "<script>window.location=\"/case/exam?rid=%d\"</script>", rid)
		}
	}
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

		var rset []Record
		if len(r.Form["pid"]) > 0 {
			pid, _ := strconv.Atoi(r.Form["pid"][0])
			for rid, rc := range rcslice {
				if rcslice[rid].PatientID == pid {
					rset = append(rset, rc)
				}
			}

		} else {
			rset = rcslice
		}

		t, _ := template.ParseFiles("static/template/caselist.gtpl")
		t.Execute(w, rset)
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
			t, err := template.ParseFiles("static/template/recorddetail.html")
			rcid, _ := strconv.Atoi(r.Form["rid"][0])
			rc := rcslice[rcid]
			err = t.Execute(w, rc)
			log.Print("err:\t", err)
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
