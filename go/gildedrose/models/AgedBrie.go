package models

type AgedBrie struct {
	SuperItem
}

func (a *AgedBrie) Update() {
	a.IncreaseQuality()
	if a.IsCaducated() {
		a.IncreaseQuality()
	}
}
