package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
)

func openbrowser() {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", "http://localhost:4001/").Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:4001/").Start()
	case "darwin":
		err = exec.Command("open", "http://localhost:4001/").Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello case!") //这个写入到w的是输出到客户端的
}

func main() {

	openbrowser()

	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":4001", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
