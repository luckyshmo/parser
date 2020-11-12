package main

import (
	"fmt"
)

func main() {
	strToArea := map[string]wArea{
		"SHOP":  Shop,
		"FIELD": Field,
	}

	tJoint := newJoint("1", strToArea["SHOP"])

	fmt.Println(tJoint)
	tJoint.setArea(Field)
	fmt.Printf("%s\n", tJoint.getInfo())
}
