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
	Date        string
}

func welcome(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	db, err := sql.Open("sqlite3", "./case.s3db")
	rows, err := db.Query("select RecordID,Patient.Name,CreateDate from Record inner join Patient on Record.PatientID=Patient.PatientID limit 10")
	checkErr(err)

	var cbarray [10]CaseBrief
	c := cbarray[0:0]

	for rows.Next() {
		var cb CaseBrief

		err = rows.Scan(&cb.CaseID, &cb.PatientName, &cb.Date)
		checkErr(err)

		c = append(c, cb)
		/*
			for _, b := range []byte(cb.PatientName) {
				fmt.Printf("%x\n", b)
			}*/
	}

	//log.Print(c)

	t, _ := template.ParseFiles("welcome.gtpl")
	t.Execute(w, c)
}

func main() {
	openbrowser("http://localhost:2273")

	//设置访问的路由
	http.HandleFunc("/", welcome)

	http.HandleFunc("/patient/search", PatientSearsh)
	http.HandleFunc("/patient/new", PatientNew)
	http.HandleFunc("/patient/update", PatientUpdate)
	http.HandleFunc("/case/new", CaseNew)
	http.HandleFunc("/case/list", CaseList)
	http.HandleFunc("/case/detail", CaseDetail)

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
