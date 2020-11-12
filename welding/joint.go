package welding

import (
	"fmt"
)

/** Joint - прдеставление сварного стыка **/
type Joint struct {
	number string
	area   Area
}

func newJoint(number string, area Area) Joint {
	return Joint{
		number: number,
		area:   area,
	}
}

func (j Joint) getInfo() string { //метод структуры //func (s structName) methodName(arg1 type, arg2 type ...) returnType {}
	return fmt.Sprintf("Стык: %s\nЗона: %s\n", j.number, j.area.name) //возможен ли nullPointer?
}

func (j *Joint) setArea(area Area) { //указатель => работаем не с копией,а по ссылке
	j.area = area
}
