package main

import (
	"fmt"

	"os"
	"strconv"
	"bufio"
)

var inputFolderPath string
var outputFolderPath string
var standartPath string
var CountVal int
var errPath = standartPath + "errlist/errList"
var e string
var numberOfRow int
var ExtentionString string

func main() {
	flagInitialize()
	realMain()
}

func realMain() {

	if errPath != "" {
		OpenedF, _ := os.Open(errPath)
		fileScanr := bufio.NewScanner(OpenedF)
		i := 0
		for fileScanr.Scan() {
			i++
			v := fileScanr.Text()
			if i == 1 {
				ExtentionString = v

			}
			if i == 2 {
				numberr, _ := strconv.Atoi(v)
				CountVal = numberr
				//fmt.Print(numberr)

			} else {
				e = v
			}

		}

		//fmt.Println(e)
	}
	//fmt.Println(e[len(e)-1:])

	returnIsFolder := checkInputOutputFolder()
	//fmt.Println(returnIsFolder)
	if returnIsFolder == true {
		fmt.Println("Tekrar Hosgeldiniz")

	} else {

		fmt.Println("Pathinizi giriniz:[girmezseniz bulundugunuz dizin kullanilacaktir]")
		fmt.Scanln(&standartPath)

		folderPathFiller()

		io := IoController{InPut: inputFolderPath, OutPut: outputFolderPath}

		io.GetFolderListOrCreateFolder(1)
		io.GetFolderListOrCreateFolder(2)

	}

	//fmt.Print(ExtentionString)
	inputLogNameArray := PathOkLogContol()

	//fmt.Println("gggggg", inputtxtNameArray)
	fc := FileCheckandScan{inputFolderPath}

	domains := fc.ExtensionControlAndScan(inputLogNameArray)

	if ExtentionString == "" {

		fmt.Println("Ozellikle istediginiz bir uzanti varsa giriniz:[ornek: com ] yoksa all yazirniz...")
		fmt.Scanln(&ExtentionString)
		WriteStringFile(ExtentionString)
	}

	if CountVal == 50 {
		fmt.Println("bir Count belirleyiniz:")
		fmt.Scanln(&CountVal)
		WriteNumberFile(CountVal)
	}
	//fmt.Print(ExtentionString)
	number, _ := strconv.Atoi(e)
	numberOfRow = number
	//if number > 1 {
	//	number = number - 1
	//}
	dm := DomainQuest{domains[numberOfRow:], CountVal, ExtentionString}

	dm.domainQueryAndWrite()

	var err = os.Remove(errPath)
	if err != nil {
		fmt.Print("dosya silinemedi")
	}

	fmt.Println("Delete Millstones")

	// i nin teker teker articagim bir yer bul
	//fl := FileCheckandScan{}

}
