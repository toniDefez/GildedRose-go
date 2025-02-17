package models

type Conjured struct {
	SuperItem
}

func (c *Conjured) Update() {
	c.DecreaseSellIn()
	c.DecreaseQualityByAmount(2)
	if c.IsCaducated() {
		c.DecreaseQualityByAmount(2)
	}
}
