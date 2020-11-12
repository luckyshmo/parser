package welding

import (
	"fmt"
)

//Joint - прдеставление сварного стыка
type Joint struct {
	Number string
	Area   Area
}

//NewJoint ...
func NewJoint(number string, area Area) Joint {
	return Joint{
		Number: number,
		Area:   area,
	}
}

//GetInfo ...
func (j Joint) GetInfo() string { //метод структуры //func (s structName) methodName(arg1 type, arg2 type ...) returnType {}
	return fmt.Sprintf("Стык: %s\nЗона: %s\n", j.Number, j.Area.Name) //возможен ли nullPointer?
}

//SetArea ...
func (j *Joint) SetArea(area Area) { //указатель => работаем не с копией,а по ссылке
	j.Area = area
}
