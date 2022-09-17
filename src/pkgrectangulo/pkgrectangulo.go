package pkgrectangulo


type Rectangulo struct {
	Base	float64
	Largo	float64
}

func (r Rectangulo) Area () float64{
	return r.Base * r.Largo
}
