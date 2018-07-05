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
	UNKNOWN  Figure = "🔵"
	SPACENUM int    = 100
)

//var marks = map[bool]string{
//	true: "✔︎", false: "❌" ,
//}
func (fig Figure) convert() string {
	switch fig {
	case OK:
		return "SUCCESS"
	case TAKEN:
		return "ERROR"
	default:
		return fmt.Sprintf("%d", string(fig))
	}
	return "bnabuyon"
}

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

