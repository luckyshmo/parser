package welding

// type wArea string //как вариант.
//Для своих типов мы также может объявлять методы с ресивером.

//Area - is welding area of joint
type Area struct {
	Name string
}

func NewArea(_name string) Area {
	return Area{
		Name: _name,
	}
}

//Shop - эстакады
var Shop = NewArea("SHOP")

//Field - это кек
var Field = NewArea("FIELD")

//NameToArea is Map str to area
var NameToArea = map[string]Area{
	"SHOP":  Shop,
	"FIELD": Field,
}
