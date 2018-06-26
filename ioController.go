package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"time"
	"strings"
)

type IoController struct {
	InPut       string
	OutPut      string
	filePath    string
	openFile    *os.File
	openFileErr error
}

func (io *IoController) setFilePath(filePath string) {
	var newPath string = filePath
	io.openFile, io.openFileErr = os.OpenFile(newPath, os.O_APPEND|os.O_WRONLY, 0644)

}

func (io IoController) appendFile(val string) {
	io.openFile.WriteString(val)
	if io.openFile == nil {
		fmt.Print("setFilePathle io.openFile i doldurmalisin")
	}

}

func (io IoController) GetFolderList(val string) ([]string, error) {
	var returnValue = []string{}

	files, err := ioutil.ReadDir(val)
	// ./ nerdeysen onu arar
	//if err != nil {
	//	//fmt.Println("Dosya Bulunamadi")
	//	//fmt.Println("soyledigin ", files)
	//}

	for _, f := range files {

		returnValue = append(returnValue, f.Name())

	}

	return returnValue, err
}


// buraci dic create eder privatetir

func (io IoController) createDic(val int) {
	//
	if val == 1 {
		os.MkdirAll(io.InPut, os.ModePerm)
	} else {
		os.MkdirAll(io.OutPut, os.ModePerm)
	}
}

func (io IoController) CreateTestFile(val string, lval string, nm string, ny string) {
	//	var returnVal bool

	if val == "yes" || val == "" {

		var file, err = os.Create("./" + lval + "/" + nm)
		if err != nil {
			fmt.Println("Dosya Olusturulamadi...")
		}
		defer file.Close()

		io.setFilePath("./" + lval + "/" + nm)
		io.appendFile(ny)

		//returnVal = true

	} else {
		fmt.Println("lutfen " + lval + " folderiniza domainlerin yazili oldugu bir dosya yerlestiriniz")
		//returnVal = false
	}
	//return returnVal
}

func (io IoController) GetFolderListOrCreateFolder(val int) []string {
	var returnVal = []string{}
	var err error
	var path string
	var pathInt int

	if val == 1 {

		path = io.InPut
		pathInt = 1

	} else {
		path = io.OutPut
		pathInt = 2

	}

	returnVal, err = io.GetFolderList(path)

	if err != nil {

		io.createDic(pathInt)
	}
	return returnVal

}
func (ty IoController) CreateAndWriteDateFile(Domain string) {
	t, _ := ty.GetFolderList(outputFolderPath)

	//for i := 0; i < len(UygunDomains); i++ {

	tn := time.Now()

	a := tn.Format("01-02-2006")

	var DateFolderName = a + ".txt"
	for _, y := range t {

		// bu splitten daha iyi time konusunda iki kere cikmasinin sebebi iki kere donmesi
		if y == DateFolderName {

			ty.setFilePath("./output/" + DateFolderName)
			ty.appendFile(Domain)

		} else {
			os.Create("./output/" + DateFolderName)
			//ty.setFilePath("./output/" + DateFolderName)
			//ty.appendFile(Domain)

		}
	}

	//}

}
func PathOkLogContol() []string {

	var returnval = []string{}

	ioa := IoController{}
	folderPathFiller()

	FolderControl, _ := ioa.GetFolderList(inputFolderPath)

	Foldercheck := 0

	for _, inputcheck := range FolderControl {

		if strings.HasPrefix(inputcheck, ".") {

		} else {

			Foldercheck++
		}

	}

	//gopre.Pre(len(inputFolderControl) == 0)
	if Foldercheck == 0 {
		//burada input folder icerisine test dosyasi sorulacak
		var answer string
		fmt.Println("Input Folderinizda herhangi bir dosya yok ornek bir dosya olusturmamizi ister misiniz? ")
		fmt.Scanln(&answer)
		ioa.CreateTestFile(answer, inputName, "tester.txt", "deneme.com\n"+"deneme1.com\n"+"deneme2.com\n"+"ojnjkdjfidjfi.com\n")

	}

	inputFolderControl2, _ := ioa.GetFolderList(inputFolderPath)

	Foldercheck2 := 0

	for _, inputcheck2 := range inputFolderControl2 {

		if strings.HasPrefix(inputcheck2, ".") {

		} else {

			// en son haline bakilsin diye actim
			//gopre.Pre(inputFolderPath)
			// pathler var demek
			//adam kullanmissa buradan devam ediyor
			returnval = append(returnval, inputcheck2)
			Foldercheck2++
		}

	}

	//virusScanProgram := filesScanner{inputFolderControl2[1:]}
	//virusScanProgram.HaydiScan()
	// txt sayisi lazim olursa kullan

	return returnval
}
