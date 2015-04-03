package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type page struct {
	Name     string
	Title    string
	Path     string
	Draft    string
	Handle   string
	Template string
	Status   int
	Behavior []behavior
	Remark   string
}

type behavior struct {
	Event  int
	Script string
	Next   int
}

var pmap map[string]page

func pageinit() {
	pmap = make(map[string]page)
	pagebyte, _ := ioutil.ReadFile("static\\template\\page.yaml")
	yaml.Unmarshal(pagebyte, &pmap)
	//log.Print(pmap)

	//d, _ := yaml.Marshal(&Page)
	//log.Printf("--- Page:\n%s\n\n", string(d))

}

func pagetestdata() {
	b1 := behavior{1, "b1 srcipt", 2}
	b2 := behavior{2, "b2 srcipt", 3}

	p1 := page{"welcome", "首页", "/", "http://www.processon.com/view/link/551cadf5e4b0f7080ab5e13d", "welcome.gtpl", "welcome", 1, []behavior{b1, b2}, "需要自动检测是否在线，显示正确的内容。\n\r需要自动检测是否有密钥对等信息，显示正确内容。"}
	p2 := page{"1.1", "四诊信息", "/record/new", "http://www.processon.com/view/link/551b95bce4b03f4b8f07caec", "welcome.gtpl", "welcome", 1, []behavior{b2, b1}, ""}

	pmap = make(map[string]page)
	pmap["/"] = p1
	pmap["/record/new"] = p2

	d, _ := yaml.Marshal(&pmap)

	//p := pages{[]page{p1, p2}}
	//d, _ := yaml.Marshal(&p)
	log.Printf("--- Page:\n%s\n\n", string(d))

	//ioutil.WriteFile("static\\template\\page.yaml", d, 0644)
}
