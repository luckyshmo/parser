package welding

// type wArea string //как вариант.
//Для своих типов мы также может объявлять методы с ресивером.

//Area - is welding area of joint
type Area struct {
	name string
}

func newArea(_name string) Area {
	return Area{
		name: _name,
	}
}

//Shop - эстакады
var Shop = newArea("SHOP")

//Field - это кек
var Field = newArea("FIELD")

//NameToArea is Map str to area
var NameToArea = map[string]Area{
	"SHOP":  Shop,
	"FIELD": Field,
}
