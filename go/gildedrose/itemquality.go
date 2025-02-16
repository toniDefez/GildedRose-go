package gildedrose

const (
	MAX_QUALITY_VALUE = 50
	MIN_QUALITY_VALUE = 0
)

type ItemQuality struct {
	Value int
}

func (i *ItemQuality) IncreaseQuality() {
	if i.Value < MAX_QUALITY_VALUE {
		i.Value = i.Value + 1
	}
}

func (i *ItemQuality) DecreaseQuality() {
	if i.Value < MIN_QUALITY_VALUE {
		i.Value = i.Value - 1
	}
}

func (i *ItemQuality) ResetQuality() {
	i.Value = MIN_QUALITY_VALUE
}

func (i *ItemQuality) DecreaseQualityByAmount(amount int) {
	temp := i.Value - amount
	if temp > MIN_QUALITY_VALUE {
		i.Value = temp
	} else {
		i.Value = MIN_QUALITY_VALUE
	}
}
