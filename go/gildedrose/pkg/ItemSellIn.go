package pkg

type ItemSellin struct {
	Value int
}

func (i *ItemSellin) DecreaseSellIn() {
	i.Value--
}

func (s ItemSellin) IsCaducated() bool {
	return s.Value <= 0
}

func (s ItemSellin) IsLessThan(days int) bool {
	return s.Value < days
}
