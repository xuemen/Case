package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
)

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

type CaseBrief struct {
	CaseID      int
	PatientName string
	CreateTime  string
}

func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("path", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	if r.Method == "GET" {
		t, _ := template.ParseFiles("static/template/register.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		if len(r.Form["pub"][0]) > 0 {
			ioutil.WriteFile("key.pub", []byte(r.Form["pub"][0]), 0644)
		}
		if len(r.Form["sec"][0]) > 0 {
			ioutil.WriteFile("key.sec", []byte(r.Form["sec"][0]), 0644)
		}

	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("path", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	db, err := sql.Open("sqlite3", "./case.v0.1.s3db")
	rows, err := db.Query("select RecordID,Patient.Name,datetime(CreateTime) from Record inner join Patient on Record.PatientID=Patient.PatientID limit 10")
	checkErr(err)

	var cbarray [10]CaseBrief
	c := cbarray[0:0]

	for rows.Next() {
		var cb CaseBrief

		err = rows.Scan(&cb.CaseID, &cb.PatientName, &cb.CreateTime)
		checkErr(err)

		c = append(c, cb)
		/*
			for _, b := range []byte(cb.PatientName) {
				fmt.Printf("%x\n", b)
			}*/
	}

	//log.Print(c)

	t, _ := template.ParseFiles("static/template/welcome.gtpl")
	t.Execute(w, c)
}

func serveFile(pattern string, filename string) {
	log.Printf("pattern:%s\tfilename:%s", pattern, filename)
	http.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, filename)
	})
}

func pagefsm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("path", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
}

func main() {
	//configdata()
	configinit()
	ret := ReadJSKey()
	if ret {
		openbrowser("http://127.0.0.1:2273")
	} else {
		openbrowser("http://127.0.0.1:2273/register")
	}

	//  yamlcleardata()
	//	yamltestdata()
	yamlinit()

	pageinit()
	log.Print(pmap)

	for k, v := range pmap {
		log.Printf("k=%v,v=%v", k, v)
		if v.Handle == "pagefsm" {
			http.HandleFunc(k, pagefsm)
		}
	}
	http.HandleFunc("/", welcome)
	http.HandleFunc("/register", register)

	/*
		//设置访问的路由
		// web pages
		http.HandleFunc("/", welcome)

		http.HandleFunc("/patient/search", PatientSearsh)
		http.HandleFunc("/patient/new", PatientNew)
		http.HandleFunc("/patient/update", PatientUpdate)
		http.HandleFunc("/case/new", CaseNew)
		http.HandleFunc("/case/list", CaseList)
		http.HandleFunc("/case/detail", CaseDetail)

	*/

	// ajax
	http.HandleFunc("/patient/info", PatientInfo)
	http.HandleFunc("/patient/brief", PatientBrief)
	http.HandleFunc("/case/info", CaseInfo)

	// static files
	http.HandleFunc("/static/", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, req.URL.Path[1:])
	})
	serveFile("/favicon.ico", "./favicon.ico")

	err := http.ListenAndServe(":2273", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
