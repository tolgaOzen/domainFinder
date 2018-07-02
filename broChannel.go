package main

import (
	"strings"
	"time"
)

type broChannel struct {
	mesaj int
}

func (a broChannel) inito() bool {
	//a := broChannel{}
	tt := make(chan int)

	go logControlInt(inputFolderPath, tt)

	go a.chanTest(tt)

	if a.chanTest(tt) == true {
		return true
	}

	if a.chanTest(tt) == false {
		return false
	}

	return false

}

func (a broChannel) chanTest(tt chan int) bool {
	exit := false

	for exit == false {
		{

			a.mesaj = <-tt

			if a.mesaj == 1 {
				//gopre.Pre(a.mesaj)
				return true
				exit = true
			}

			if a.mesaj == 0 {
				//gopre.Pre(a.mesaj)
				return false

			}

			//fmt.Println(a.mesaj)
			time.Sleep(time.Second)

		}
	}
	return false
}

func logControlInt(path string, tt chan int) {
	io := IoController{}

	for {
		{
			a, _ := io.GetFolderList(path)
			//gopre.Pre(path)
			total := 0
			for _, b := range a {
				//gopre.Pre(b)
				if strings.HasPrefix(b, ".") {

				} else {
					total++

				}
			}

			tt <- total

			time.Sleep(time.Second)

		}
	}

}
