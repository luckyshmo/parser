package main

import (
	"fmt"
	"parser/welding"
)

func main() {

	tJoint := welding.NewJoint("1", welding.NameToArea["SHOP"])

	fmt.Println(tJoint)
	tJoint.SetArea(welding.Field)
	fmt.Printf("%s\n", tJoint.GetInfo())
}
