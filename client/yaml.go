package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Index struct {
	MaxPatientID int
	MaxRecordID  int
}

var index Index
var pslice []Patient
var rslice []CaseDetailData

func yamlinit() {
	indexbyte, _ := ioutil.ReadFile("data\\index.yaml")
	yaml.Unmarshal(indexbyte, &index)
	//log.Print(index)

	//d, _ := yaml.Marshal(&index)
	//log.Printf("--- index:\n%s\n\n", string(d))

	var p Patient
	var r CaseDetailData
	var filename string
	var pbyte, rbyte []byte
	for pid := 0; pid <= index.MaxPatientID; pid++ {
		filename = fmt.Sprintf("data\\patient\\%d.yaml", pid)
		pbyte, _ = ioutil.ReadFile(filename)
		yaml.Unmarshal(pbyte, &p)

		pslice = append(pslice, p)
	}

	for rid := 0; rid <= index.MaxRecordID; rid++ {
		filename = fmt.Sprintf("data\\record\\%d.yaml", rid)
		rbyte, _ = ioutil.ReadFile(filename)
		yaml.Unmarshal(rbyte, &r)

		rslice = append(rslice, r)
	}
}

func yamltestdata() {
	testp := Patient{0, "test patient", "male", "19750322", "test address", "英文缩写：PMH 英文：Past Medical History", "英文缩写：FMH 英文：Family Medical History", "英文：Allergies"}
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
		r := CaseBrief{rslice[rid].RecordID, pslice[rslice[rid].PatientID].Name, rslice[rid].CreateTime}
		ret = append(ret, r)
	}

	return ret
}
