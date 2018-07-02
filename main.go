package main

import "fmt"

var inputFolderPath string
var outputFolderPath string
var standartPath string

func main() {
	flagInitialize()
	realMain()
}
func realMain() {

	returnIsFolder := checkInputOutputFolder()
	//fmt.Println(returnIsFolder)
	if returnIsFolder == true {
		fmt.Println("Tekrar Hosgeldiniz")
		//folderPathFiller()
		//PathlerVarInputtaKacTxtvar()
		//hasTxt :=
		// burada inputta kac tane txt var o yazar

		//gopre.Pre(hasTxt)

	} else {

		fmt.Println("Pathinizi giriniz:[girmezseniz bulundugunuz dizin kullanilacaktir]")
		fmt.Scanln(&standartPath)

		folderPathFiller()

		io := IoController{InPut: inputFolderPath, OutPut: outputFolderPath}

		io.GetFolderListOrCreateFolder(1)
		io.GetFolderListOrCreateFolder(2)

	}

	inputLogNameArray := PathOkLogContol()
	//fmt.Println("gggggg", inputtxtNameArray)
	fc := FileCheckandScan{inputFolderPath}
	domains := fc.ExtensionControlAndScan(inputLogNameArray)

	var ExtentionString string
	fmt.Println("Ozellikle istediginiz bir uzanti varsa giriniz:[ornek: com ] yoksa all yazirniz...")
	fmt.Scanln(&ExtentionString)

	if CountVal == 50 {
		fmt.Println("bir Count belirleyiniz:")
		fmt.Scanln(&CountVal)

	}

	dm := DomainQuest{domains, CountVal, ExtentionString}

	dm.domainQueryAndWrite()

}
