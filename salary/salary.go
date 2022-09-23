package salary


type Salary float64
func (s Salary) Total() Total {
	return Total(s)
}

type Total float64
func (t Total) Hra() Hra{
	t += t*0.3
	return Hra(t)
}
func (t Total) Salary() Salary{
	t -= t*0.1
	return Salary(t)
}

type Hra float64
func (h Hra) Basic() Basic {
   h += h * 0.3   // 30% HRA Addition
  return Basic(h)
}
func (h Hra) Total() Total {   
	return Total(h)
}

type Basic float64
func (b Basic) Total() Total {
	return Total(b)
}
  