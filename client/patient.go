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
	"time"
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
		var ret []PatientSearchResult
		HasResult := false
		for pid, v := range pslice {
			var psr PatientSearchResult
			log.Printf("pid=%d\tv.PatientID=%d", pid, v.PatientID)
			if (len(r.Form["id"][0]) > 0) && (string(v.PatientID) != r.Form["id"][0]) {
				log.Printf("r.Form[\"id\"][0]=%s", r.Form["id"][0])
				continue
			}

			if (len(r.Form["name"][0]) > 0) && (v.Name != r.Form["name"][0]) {
				log.Printf("r.Form[\"name\"][0]=%s", r.Form["name"][0])
				continue
			}

			if (r.Form["sex"][0] != "All") && (v.Sex != r.Form["sex"][0]) {
				log.Printf("r.Form[\"sex\"][0]=%s", r.Form["sex"][0])
				continue
			}
			if (len(r.Form["DOB"][0]) > 0) && (v.DOB != r.Form["DOB"][0]) {
				log.Printf("r.Form[\"DOB\"][0]=%s", r.Form["DOB"][0])
				continue
			}

			rtime, _ := time.Parse("1900-01-01 10:30:00", "1900-01-01 10:30:00")
			for rid := index.MaxRecordID; rid >= 0; rid-- {
				if rslice[rid].PatientID == v.PatientID {
					rtime, _ = time.Parse("2015-01-01 10:30:00", rslice[rid].CreateTime)
					psr.CreateTime = rslice[rid].CreateTime
					psr.Diag = rslice[rid].Diag
					break
				}
			}

			if r.Form["time"][0] != "All" {

				if (r.Form["time"][0] == "24h") && (time.Since(rtime).Hours() >= 24.0) {
					continue
				}
				if (r.Form["time"][0] == "7d") && (time.Since(rtime).Hours() >= 7*24.0) {
					continue
				}
			}

			for rid := index.MaxRecordID; rid >= 0; rid-- {
				if rslice[rid].PatientID == v.PatientID {
					rtime, _ = time.Parse("2015-01-01 10:30:00", rslice[rid].CreateTime)
					break
				}
			}

			psr.Address = v.Address
			psr.DOB = v.DOB
			psr.Name = v.Name
			psr.PatientID = v.PatientID
			psr.Sex = v.Sex

			ret = append(ret, psr)
			HasResult = true
		}

		t, _ := template.ParseFiles("static/template/patientsearch.gtpl")
		if HasResult {
			t.Execute(w, ret)
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
		pslice = append(pslice, p)
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
