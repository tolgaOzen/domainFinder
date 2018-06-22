package main

import (
	"strings"
	"fmt"
	"github.com/tolgaOzen/matrixPrint"
)

type Figure string

const (
	OK       Figure = "✔︎"
	TAKEN    Figure = "❌"
	SPACENUM int    = 100
)

//var marks = map[bool]string{
//	true: "✔︎", false: "❌" ,
//}

type Write struct {
	figure Figure
}

func (write Write) spaceWrite(row string) {

	RowLen := len(row)

	Sum := 1 + RowLen

	SpaceNum := strings.Repeat(" ", SPACENUM-Sum)

	matrix.MPrint(row)
	fmt.Print(SpaceNum)
	fmt.Print(write.figure + "\n")
	//matrix.MPrint(write.figure + "\n")

}
