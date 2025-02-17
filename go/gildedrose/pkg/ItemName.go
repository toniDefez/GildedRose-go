package pkg

const (
	AGED_BRIE = "Aged Brie"
	CONJURED  = "Conjured"
	BACKSTAGE = "Backstage passes to a TAFKAL80ETC concert"
	SULFURAS  = "Sulfuras, Hand of Ragnaros"
)

type ItemName struct {
	Name string
}

func (i ItemName) IsAgedBrie() bool {
	return i.Name == AGED_BRIE
}

func (i ItemName) IsConjured() bool {
	return i.Name == CONJURED
}

func (i ItemName) IsBackstagePasses() bool {
	return i.Name == BACKSTAGE
}

func (i ItemName) IsSulfuras() bool {
	return i.Name == SULFURAS
}
