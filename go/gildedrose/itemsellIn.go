package gildedrose

type ItemSellin struct {
	Value int
}

func (i *ItemSellin) DecreaseSellIn() {
	i.Value--
}

func (s ItemSellin) IsCaducated() bool {
	return s.Value < 0
}

func (s ItemSellin) isLessThan(days int) bool {
	return s.Value < days
}
