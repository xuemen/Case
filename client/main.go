package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
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

	yamlcleardata()
	//	yamltestdata()
	yamlinit()

	pageinit()
	log.Print(pmap)

	for k, v := range pmap {
		log.Printf("k=%v,v=%v", k, v)
		http.HandleFunc(k, pagefsm)
	}

	openbrowser("http://127.0.0.1:2273")

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
