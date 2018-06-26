package main

import (
	"flag"
)

var help string
var CountVal int

func flagInitialize() {
	//var wg = new(sync.WaitGroup)
	//wg.Add(3)
	//input output kloser isimleri alincak
	//alinan kloser adlari kontrol edilecek


	// standartpath hale getir
	flag.StringVar(&inputFolderPath, "i", "", "input Path i  dogru bir sekilde girmelisin")
	flag.StringVar(&outputFolderPath, "o", "", "output Path i dogru bir sekilde girmelisin")
	flag.IntVar(&CountVal, "c", 50, "bir count girmelisin domainler belirledigin counta gore listene eklenir belirlemezsen 20 kabul edilir")
	flag.StringVar(&help, "help", "", "-i komutundan sonra bir input path i girmelisin \n -o komutundan sonra bir output path girmelisin \n -c komutundan sonra bir count girmelisin domainler belirledigin counta gore listene eklenir belirlemezsen 20 kabul edilir ")
	flag.Parse()
	//calistigim klasor icerinsinde input ve output klasorleri var mi kontrol et
}

func checkInputOutputFolder() bool {
	var returnVal bool = false

	iop := IoController{}
	folder, _ := iop.GetFolderList("./")
	//gopre.Pre(folder)
	var err = 0
	for _, fol := range folder {

		if fol != inputName {
			err++
			//fmt.Println("input name buldum", fol)

		}
		if fol != outputName {
			err++
			//fmt.Println("output u buldum", fol)
		}
	}

	//fmt.Println(err, len(folder)*2)
	var totalCount = len(folder) * 2

	if err == totalCount-2 {

		//matrix.MPrint("true")
		returnVal = true
	} else {
		//matrix.MPrint("false")
		returnVal = false
	}

	return returnVal
}

func folderPathFiller() {

	if standartPath == "" {

		standartPath = "./"
	}
	//gopre.Pre(standartPath)

	if inputFolderPath == "" {

		inputFolderPath = standartPath + inputName

	}
	if outputFolderPath == "" {

		outputFolderPath = standartPath + outputName
	}

}