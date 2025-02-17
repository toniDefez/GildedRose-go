package models

const (
	BACKSTAGE_PASSES_DOUBLE_INCREASE = 10
	BACKSTAGE_PASSES_TRIPLE_INCREASE = 5
	BACKSTAGE_PASSES_RESET           = 0
)

type Backstage struct {
	SuperItem
}

func (b *Backstage) Update() {
	b.DecreaseSellIn()
	b.IncreaseQuality()

	if b.HasToBeSoldInLessThan(BACKSTAGE_PASSES_DOUBLE_INCREASE) {
		b.IncreaseQuality()
	}
	if b.HasToBeSoldInLessThan(BACKSTAGE_PASSES_TRIPLE_INCREASE) {
		b.IncreaseQuality()
	}
	if b.HasToBeSoldInLessThan(BACKSTAGE_PASSES_RESET) {
		b.ResetQuality()
	}
}
