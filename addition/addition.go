package addition

type Addition int 
type Multiply int

func (a *Addition) Twice(){
	*a = Addition(*a+*a)
}

func (m *Multiply) Twice(){
	*m = Multiply(*m * 2)
}