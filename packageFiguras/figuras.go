package packagefiguras

type Cuadrado struct{
	Nombre string
	Base float64
}

type Rectangulo struct{
	Nombre string
	Base float64
	Altura float64
}

type Triangulo struct{
	Nombre string
	Lado float64
}

func (c Cuadrado) Area() float64{
	return c.Base * c.Base
}

func (r Rectangulo) Area() float64{
	return r.Base * r.Altura
}

func (t Triangulo) Area() float64{
	return t.Lado * t.Lado * t.Lado
}