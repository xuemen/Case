package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type Index struct {
	MaxPatientID int
	MaxRecordID  int
	MaxRecptID   int
}

var index Index
var pslice []Patient
var rcslice []Record
var rtslice []Recpt

func yamlinit() {
	indexbyte, _ := ioutil.ReadFile("data\\index.yaml")
	yaml.Unmarshal(indexbyte, &index)
	//log.Print(index)

	//d, _ := yaml.Marshal(&index)
	//log.Printf("--- index:\n%s\n\n", string(d))

	var p Patient
	var rc Record
	var rt Recpt
	var filename string
	var pbyte, rcbyte, rtbyte []byte
	for pid := 0; pid <= index.MaxPatientID; pid++ {
		filename = fmt.Sprintf("data\\patient\\%d.yaml", pid)
		pbyte, _ = ioutil.ReadFile(filename)
		yaml.Unmarshal(pbyte, &p)

		pslice = append(pslice, p)
	}

	for rcid := 0; rcid <= index.MaxRecordID; rcid++ {
		filename = fmt.Sprintf("data\\record\\%d.yaml", rcid)
		rcbyte, _ = ioutil.ReadFile(filename)
		yaml.Unmarshal(rcbyte, &rc)

		rcslice = append(rcslice, rc)
	}

	for rtid := 0; rtid <= index.MaxRecptID; rtid++ {
		filename = fmt.Sprintf("data\\recpt\\%d.yaml", rtid)
		rtbyte, _ = ioutil.ReadFile(filename)
		yaml.Unmarshal(rtbyte, &rt)

		rtslice = append(rtslice, rt)
	}
}

func yamltestdata() {
	testp := Patient{0, "test patient", 0, "男", "19750322", 54, "未婚", "软件工程师", "中华人民共和国", "汉族", "广西梧州", "13910911670", "北京市海淀区车公庄西路35号院北工大留创院121室", "100044", "201504010101000", "英文缩写：PMH 英文：Past Medical History", "英文缩写：FMH 英文：Family Medical History", "英文：Allergies"}
	var d []byte
	var filename string

	for pid := 0; pid < 10; pid++ {
		testp.PatientID = pid
		d, _ = yaml.Marshal(&testp)
		//log.Printf("--- testp:\n%s\n\n", string(d))

		filename = fmt.Sprintf("data\\patient\\%d.yaml", pid)
		ioutil.WriteFile(filename, d, 0644)
	}
	index.MaxPatientID = 9

	testr := CaseDetailData{0, 0, "英文：Main Complaint", "ExamReport", "Diag", "DRR", "英文：Presciption", "", "20150401", "true"}
	for rid := 0; rid < 20; rid++ {
		testr.RecordID = rid
		testr.PatientID = rid / 2
		d, _ = yaml.Marshal(&testr)
		//log.Printf("--- testr:\n%s\n\n", string(d))

		filename = fmt.Sprintf("data\\record\\%d.yaml", rid)
		ioutil.WriteFile(filename, d, 0644)
	}
	index.MaxRecordID = 19

	d, _ = yaml.Marshal(&index)
	ioutil.WriteFile("data\\index.yaml", d, 0644)
}

func getnewrid() int {
	index.MaxRecordID = index.MaxRecordID + 1
	rid := index.MaxRecordID

	d, _ := yaml.Marshal(&index)
	ioutil.WriteFile("data\\index.yaml", d, 0644)

	return rid
}

func saverecord(rid int, rc Record) {
	log.Printf("saverecord:\nrid=%d\ndata=%v", rid, rc)

	d, _ := yaml.Marshal(&rc)

	filename := fmt.Sprintf("data\\record\\%d.yaml", rid)
	ioutil.WriteFile(filename, d, 0644)

	log.Printf("len(rcslice)=%d", len(rcslice))
	if rid > len(rcslice) {
		rcslice = append(rcslice, rc)
	} else {
		rcslice[rid] = rc
	}

}

func yamlcleardata() {
	indexbyte, _ := ioutil.ReadFile("data\\index.yaml")
	yaml.Unmarshal(indexbyte, &index)

	var filename string
	for pid := 0; pid <= index.MaxPatientID; pid++ {
		filename = fmt.Sprintf("data\\patient\\%d.yaml", pid)
		os.Remove(filename)
	}
	index.MaxPatientID = 0

	for rid := 0; rid <= index.MaxRecordID; rid++ {
		filename = fmt.Sprintf("data\\record\\%d.yaml", rid)
		os.Remove(filename)
	}
	index.MaxRecordID = 0

	d, _ := yaml.Marshal(&index)
	ioutil.WriteFile("data\\index.yaml", d, 0644)
}

type config struct {
	RealName      string
	UserName      string
	Email         string
	CellPhone     string
	Password      string
	KeyPassphrase string
	Org           string
	Statement     string
}

var conf config

func configinit() {
	confbyte, _ := ioutil.ReadFile("config.yaml")
	yaml.Unmarshal(confbyte, &conf)

	log.Printf("config.yaml read... username:%s", conf.UserName)
}

func configdata() {
	testconf := config{"黄勇刚", "huangyg", "huangyg@xuemen.com", "13910911670", "12344", "", "北京学门科技有限公司", "..."}

	d, _ := yaml.Marshal(&testconf)
	ioutil.WriteFile("config.yaml", d, 0644)
}

func readCaseBrief(n int) []CaseBrief {
	var ret []CaseBrief

	MinRid := 0
	if index.MaxRecordID > n {
		MinRid = index.MaxRecordID - n
	}
	for rid := index.MaxRecordID; rid >= MinRid; rid-- {
		r := CaseBrief{rcslice[rid].RecordID, pslice[rcslice[rid].PatientID].Name, rcslice[rid].CreateTime}
		ret = append(ret, r)
	}

	return ret
}

var chinc map[string]string // [医院编码]疾病名称
var chinj map[string]string // [病名缩写]医院编码

func chineseillnameiit() {
	chinc = make(map[string]string)
	chinj = make(map[string]string)

	chincbyte, _ := ioutil.ReadFile("data\\init\\chinc.yaml")
	yaml.Unmarshal(chincbyte, &chinc)
	//log.Print(chinc)

	a := pinyin.NewArgs()
	a.Style = pinyin.FirstLetter

	reg := regexp.MustCompile(`[（][^（）]+[）]`)

	for c, name := range chinc {
		name = reg.ReplaceAllString(name, "")
		name = strings.Replace(name, " ", "", -1)
		name = strings.Replace(name, "*", "", -1)
		name = strings.Replace(name, "、", "", -1)
		name = strings.Replace(name, "》", "", -1)
		name = strings.Replace(name, "?", "", -1)

		p := pinyin.LazyPinyin(name, a)
		j := strings.Join(p, "")
		_, ok := chinj[j]
		if ok {
			//同音病名处理
		}
		chinc[c] = name
		chinj[j] = c
		//log.Printf("%s\t%s\t%s", c, j, name)
	}
	d, _ := yaml.Marshal(&chinc)
	ioutil.WriteFile("data\\init\\chinc.yaml", d, 0644)
	d, _ = yaml.Marshal(&chinj)
	ioutil.WriteFile("data\\init\\chinj.yaml", d, 0644)
}
