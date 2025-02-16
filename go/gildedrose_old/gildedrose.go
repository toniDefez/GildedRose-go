package gildedrose_old

type Item struct {
	Name            string
	SellIn, Quality int
}

const (
	AGED_BRIE                        = "Aged Brie"
	CONJURED                         = "Conjured"
	BACKSTAGE                        = "Backstage passes to a TAFKAL80ETC concert"
	SULFURAS                         = "Sulfuras, Hand of Ragnaros"
	MAX_QUALITY                      = 50
	MIN_QUALITY                      = 0
	BACKSTAGE_PASSES_DOUBLE_INCREASE = 10
	BACKSTAGE_PASSES_TRIPLE_INCREASE = 5
	BACKSTAGE_PASSES_RESET           = 0
	STEP_DOWN_QUALITY_DEFAULT        = 1
	STEP_DOWN_QUALITY_CONJURED       = 2
)

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		currentItem := items[i]
		switch currentItem.Name {
		case AGED_BRIE:
			items[i] = UpdateBrie(*currentItem)
		case CONJURED:
			items[i] = UpdateConjured(*currentItem)
		case BACKSTAGE:
			items[i] = UpdateBackStage(*currentItem)
		case SULFURAS:
			items[i] = currentItem
		default:
			items[i] = UpdateDefault(*currentItem)
		}
	}

}

func DecreaseSellIn(currentItem Item) *Item {
	currentItem.SellIn--
	return &currentItem
}

func IncreaseQuality(currentItem Item) *Item {
	if currentItem.Quality < MAX_QUALITY {
		currentItem.Quality++
	}
	return &currentItem
}

func ResetQuality(currentItem Item) *Item {
	currentItem.Quality = MIN_QUALITY
	return &currentItem
}

func DecreaseQuality(currentItem Item, step int) *Item {
	if currentItem.Quality > MIN_QUALITY {
		currentItem.Quality = currentItem.Quality - step
	}
	return &currentItem
}

func UpdateBrie(currentItem Item) *Item {
	currentItem = *DecreaseSellIn(currentItem)
	currentItem = *IncreaseQuality(currentItem)
	if currentItem.SellIn <= 0 {
		currentItem = *IncreaseQuality(currentItem)
	}
	return &currentItem
}

func UpdateConjured(currentItem Item) *Item {
	currentItem = *DecreaseSellIn(currentItem)
	currentItem = *DecreaseQuality(currentItem, 2)
	if currentItem.SellIn < 0 {
		currentItem = *DecreaseQuality(currentItem, 2)
	}

	if currentItem.Quality < MIN_QUALITY {
		currentItem.Quality = MIN_QUALITY
	}
	return &currentItem
}

func UpdateBackStage(currentItem Item) *Item {
	currentItem = *DecreaseSellIn(currentItem)
	currentItem = *IncreaseQuality(currentItem)

	if currentItem.SellIn < BACKSTAGE_PASSES_DOUBLE_INCREASE {
		currentItem = *IncreaseQuality(currentItem)
	}
	if currentItem.SellIn < BACKSTAGE_PASSES_TRIPLE_INCREASE {
		currentItem = *IncreaseQuality(currentItem)
	}
	if currentItem.SellIn < BACKSTAGE_PASSES_RESET {
		currentItem = *ResetQuality(currentItem)
	}
	return &currentItem

}

func UpdateDefault(currentItem Item) *Item {
	currentItem = *DecreaseSellIn(currentItem)
	currentItem = *DecreaseQuality(currentItem, STEP_DOWN_QUALITY_DEFAULT)
	if currentItem.SellIn < 0 {
		currentItem = *DecreaseQuality(currentItem, STEP_DOWN_QUALITY_DEFAULT)
	}

	if currentItem.Quality < MIN_QUALITY {
		currentItem.Quality = MIN_QUALITY
	}
	return &currentItem
}
