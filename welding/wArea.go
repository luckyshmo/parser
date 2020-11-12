package welding

// type wArea string //как вариант.
//Для своих типов мы также может объявлять методы с ресивером.

type Area struct {
	name string
}

func newArea(_name string) Area {
	return Area{
		name: _name,
	}
}

var Shop = newArea("SHOP")
var Field = newArea("FIELD")
