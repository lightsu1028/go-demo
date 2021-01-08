package concurrency

import (
	"fmt"
	"testing"
)

type Class struct {
	Stu []*Stu
}

type Stu struct {
	Name string
}

var stMap map[string]Stu

func TestSprintf(t *testing.T) {
	//var f float64=1.34
	////var fStr string ="1.34"
	//r := fmt.Sprintf("%f%%", f)
	//
	//fmt.Println(r)
	stu := stMap["baikang"]

	fmt.Println(stu.Name)
}
