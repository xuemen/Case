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
var parray [1000]Patient
var rarray [10000]Case

func yamlinit() {
	indexbyte, _ := ioutil.ReadFile("data\\index.yaml")
	yaml.Unmarshal(indexbyte, &index)
	log.Print(index)

	d, _ := yaml.Marshal(&index)
	log.Printf("--- index:\n%s\n\n", string(d))

	var p Patient
	var r Case
	var filename string
	var pbyte, rbyte []byte
	for pid := 0; pid <= index.MaxPatientID; pid++ {
		filename = fmt.Sprintf("data\\patient\\%d.yaml", pid)
		pbyte, _ = ioutil.ReadFile(filename)
		yaml.Unmarshal(pbyte, &p)

		parray[pid] = p
	}

	for rid := 0; rid <= index.MaxRecordID; rid++ {
		filename = fmt.Sprintf("data\\record\\%d.yaml", rid)
		rbyte, _ = ioutil.ReadFile(filename)
		yaml.Unmarshal(rbyte, &r)

		rarray[rid] = r
	}
}

func yamltestdata() {
	testp := Patient{0, "test patient", "male", "19750322", "test address", "英文缩写：PMH 英文：Past Medical History", "英文缩写：FMH 英文：Family Medical History", "英文：Allergies"}
	var d []byte
	var filename string

	for pid := 0; pid < 10; pid++ {
		testp.PatientID = pid
		d, _ = yaml.Marshal(&testp)
		log.Printf("--- testp:\n%s\n\n", string(d))

		filename = fmt.Sprintf("data\\patient\\%d.yaml", pid)
		ioutil.WriteFile(filename, d, 0644)
	}
	index.MaxPatientID = 9

	testr := Case{0, "英文：Main Complaint", "ExamReport", "Diag", "DRR", "英文：Presciption", "20150401"}
	for rid := 0; rid < 20; rid++ {
		testr.RecordID = rid
		d, _ = yaml.Marshal(&testr)
		log.Printf("--- testr:\n%s\n\n", string(d))

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
