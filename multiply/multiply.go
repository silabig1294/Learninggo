package multiply

type Multiply struct {
	Num int
}

func (m *Multiply) Twice(n int) {
	m.Num = 2*n	

}

func (m Multiply) Display() int{ 
	return m.Num
}

