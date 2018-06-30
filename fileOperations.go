package main

import (
	"strings"
	"os"
	"fmt"
	"bufio"
)

//type filesScanner struct {
//	folderList []string
//}


func DomainCheck(e string) bool {

	if !strings.ContainsAny(e, "@?") && strings.ContainsAny(e, ".") &&
		!strings.HasPrefix(e, ".") && !strings.HasSuffix(e, ".") &&
		!strings.ContainsAny(e, "\t") && strings.Count(e, " ") == 0 {
		return true
	}
	return false
}

type FileCheckandScan struct {
	//extension strings
	path string
}

func (fc FileCheckandScan) ExtensionControlAndScan(txtNames []string) []string {

	var returnVal = []string{}

	for _, txtName := range txtNames {

		txtName = fc.path + "/" + txtName
		b := fc.txtScanner(txtName)

		for _, a := range b {

			returnVal = append(returnVal, a)
		}

	}
	return returnVal
}

// buraya txt name leri vericem oda bana icindekileri array donucek
func (fc FileCheckandScan) txtScanner(val string) []string {

	var returnVal = []string{}

	OpenedFile, error := os.Open(val)
	if error != nil {
		fmt.Println("Dosya Acilmadi")
	}
	defer OpenedFile.Close()
	fileScanner := bufio.NewScanner(OpenedFile)

	for fileScanner.Scan() {

		b := fileScanner.Text()

		a := strings.Split(b, " ")

		for _, c := range a {

			e := strings.TrimSpace(c)

			if DomainCheck(e) == true {
				//UGUR ABIYE SOR BIR BOSLUK VARSA ARAYA \t KOYUYO NEDENINI ANLAMADIM
				//fmt.Println("=====", e)
				returnVal = append(returnVal, e)

			}

		}

	}
	return returnVal

}

//func (fs filesScanner) HaydiScan() int {
//	var scanCount = 0
//
//	for _, scanRow := range fs.folderList {
//		t := strings.TrimSpace(scanRow)
//		e := strings.HasSuffix(t, ".txt")
//		if e == true {
//			scanCount++
//		}
//
//	}
//
//	return scanCount
//}
