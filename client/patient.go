package main

import (
	"database/sql"
	"fmt"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type PID struct {
	PatientID string
}

type Patient struct {
	PatientID   int
	Name        string
	SCBC        int
	Sex         string
	DOB         string
	Weight      float64 //kg
	Marital     string
	Career      string
	Nationality string
	Race        string
	POB         string
	Phone       string
	Address     string
	Postcode    string
	ServiceTime string
	PMH         string
	FMH         string
	Allergies   string
}

//testp := Patient{0, "test patient","0", "male", "19750322",54,"单身","软件工程师","中华人民共和国","汉族","广西梧州","13910911670", "北京市海淀区车公庄西路35号院北工大留创院121室", "100044","201504010101000"}

type PatientSearchResult struct {
	PatientID  int
	Name       string
	Sex        string
	DOB        string
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
		t, _ := template.ParseFiles("static/template/patientsearch.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		where := false
		wherestr := ""
		sqlstr := "select Patient.PatientID,Patient.Name,Patient.Sex,Patient.DOB,Patient.Address,ifnull(datetime(CreateTime),'未就诊'),ifnull(Diag,'未填写诊断') from Patient left join Record on Patient.PatientID = Record.PatientID %s group by Patient.PatientID limit 100"
		if len(r.Form["id"][0]) > 0 {
			if where {
				wherestr = fmt.Sprintf("%s and Patient.PatientID=%s", wherestr, r.Form["id"][0])
			} else {
				wherestr = fmt.Sprintf("%s Patient.PatientID=%s", wherestr, r.Form["id"][0])
			}
			where = true
			log.Print("wherestr:\t", wherestr)
		}

		if len(r.Form["name"][0]) > 0 {
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

		if len(r.Form["DOB"][0]) > 0 {
			if where {
				wherestr = fmt.Sprintf("%s and Patient.DOB=\"%s\"", wherestr, r.Form["DOB"][0])
			} else {
				wherestr = fmt.Sprintf("%s Patient.DOB=\"%s\"", wherestr, r.Form["DOB"][0])
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
		defer db.Close()
		rows, err := db.Query(sqlstr)
		checkErr(err)

		log.Print(rows)

		var parray [100]PatientSearchResult
		c := parray[0:0]

		HasResult := false

		for rows.Next() {
			var p PatientSearchResult

			err = rows.Scan(&p.PatientID, &p.Name, &p.Sex, &p.DOB, &p.Address, &p.CreateTime, &p.Diag)
			checkErr(err)

			c = append(c, p)
			HasResult = true
			log.Print("c:\t", c)
			log.Print("p:\t", p)
		}

		t, _ := template.ParseFiles("static/template/patientsearch.gtpl")
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
		t, _ := template.ParseFiles("static/template/patientnew.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		name := r.Form["name"][0]
		scbc, _ := strconv.Atoi(r.Form["scbc"][0])
		sex := r.Form["sex"][0]
		DOB := r.Form["DOB"][0]
		weight, _ := strconv.ParseFloat(r.Form["weight"][0], 32)
		marital := r.Form["marital"][0]
		career := r.Form["career"][0]
		nationality := r.Form["nationality"][0]
		race := r.Form["race"][0]
		POB := r.Form["POB"][0]
		phone := r.Form["phone"][0]
		address := r.Form["address"][0]
		postcode := r.Form["postcode"][0]
		servicetime := r.Form["servicetime"][0]
		PMH := r.Form["PMH"][0]
		FMH := r.Form["FMH"][0]
		allergies := r.Form["allergies"][0]

		index.MaxPatientID++
		d, _ := yaml.Marshal(&index)
		ioutil.WriteFile("data\\index.yaml", d, 0644)
		log.Printf("index.yaml update:\n%s\n", d)

		filename := fmt.Sprintf("data\\patient\\%d.yaml", index.MaxPatientID)
		p := Patient{index.MaxPatientID, name, scbc, sex, DOB, weight, marital, career, nationality, race, POB, phone, address, postcode, servicetime, PMH, FMH, allergies}
		d, _ = yaml.Marshal(&p)
		ioutil.WriteFile(filename, d, 0644)
		log.Printf("patient new...patientid:%d\n%s", index.MaxPatientID, d)

		fmt.Fprintf(w, "<script>window.location=\"/case/new?pid=%d\"</script>", index.MaxPatientID)
	}

}

func PatientUpdate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("Action:", r.Form["action"])
	fmt.Println("ID:", r.Form["id"])
	fmt.Println("Name:", r.Form["name"])
	fmt.Println("Sex:", r.Form["sex"])
	fmt.Println("DOB:", r.Form["DOB"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	if r.Method == "GET" {
		t, _ := template.ParseFiles("static/template/patientupdate.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {

	}

}

func PatientInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("path", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	if r.Method == "GET" {
		if len(r.Form["pid"]) > 0 {
			sqlstr := fmt.Sprintf("select PatientID,Name,Sex,ifnull(date(DOB),\"未填写生日\"),ifnull(Address,\"未填写地址\"),ifnull(PMH,\"未填写既往病史\"),ifnull(FMH,\"未填写家族病史\"),ifnull(Allergies,\"未填写过敏史\") from Patient where PatientID = %s", r.Form["pid"][0])
			log.Print("sqlstr:\t", sqlstr)

			db, err := sql.Open("sqlite3", "./case.v0.1.s3db")
			defer db.Close()

			rows, err := db.Query(sqlstr)
			checkErr(err)

			//log.Print(rows)
			var p Patient
			for rows.Next() {
				err = rows.Scan(&p.PatientID, &p.Name, &p.Sex, &p.DOB, &p.Address)
				checkErr(err)
			}
			log.Print("p:\t", p)
			t, _ := template.ParseFiles("static/template/patientinfo.gtpl")
			err = t.Execute(w, p)
			log.Print("err:\t", err)
		} else {
			fmt.Fprint(w, "未指定病人编号。")
		}
	}
}

func PatientBrief(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("path", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	if r.Method == "GET" {
		if len(r.Form["pid"]) > 0 {
			sqlstr := fmt.Sprintf("select PatientID,Name,Sex,ifnull(date(DOB),\"未填写生日\"),ifnull(Address,\"未填写地址\") from Patient where PatientID = %s", r.Form["pid"][0])
			log.Print("sqlstr:\t", sqlstr)

			db, err := sql.Open("sqlite3", "./case.v0.1.s3db")
			defer db.Close()

			rows, err := db.Query(sqlstr)
			checkErr(err)

			//log.Print(rows)
			var p Patient
			for rows.Next() {
				err = rows.Scan(&p.PatientID, &p.Name, &p.Sex, &p.DOB, &p.Address)
				checkErr(err)
			}
			log.Print("p:\t", p)
			t, _ := template.ParseFiles("static/template/patientbrief.gtpl")
			err = t.Execute(w, p)
			log.Print("err:\t", err)
		} else {
			fmt.Fprint(w, "未指定病人编号。")
		}
	}
}
