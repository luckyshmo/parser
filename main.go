package main

import (
	"fmt"
)

func main() {

	tJoint := newJoint("1", NameToArea["SHOP"])

	fmt.Println(tJoint)
	tJoint.setArea(Field)
	fmt.Printf("%s\n", tJoint.getInfo())
}
