package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		currentItem := items[i]

		switch currentItem.Name {
		case "Aged Brie":
			items[i] = UpdateBrie(*currentItem)
		case "Conjured":
			items[i] = UpdateConjured(*currentItem)
		case "Backstage passes to a TAFKAL80ETC concert":
			items[i] = UpdateBackStage(*currentItem)
		case "Sulfuras, Hand of Ragnaros":
			items[i] = currentItem
		default:
			items[i] = UpdateDefault(*currentItem)
		}
	}

}

func UpdateBrie(currentItem Item) *Item {
	currentItem.SellIn--
	if currentItem.SellIn <= 0 {
		tempQuality := currentItem.Quality + 2
		if tempQuality > 50 {
			currentItem.Quality = 50
		} else {
			currentItem.Quality = tempQuality
		}
	} else {
		currentItem.Quality++
	}
	return &currentItem
}

func UpdateConjured(currentItem Item) *Item {
	currentItem.SellIn--

	if currentItem.SellIn < 0 {
		tempQuality := currentItem.Quality - 4
		if tempQuality < 0 {
			currentItem.Quality = 0
		} else {
			currentItem.Quality = tempQuality
		}
	} else {
		tempQuality := currentItem.Quality - 2
		if tempQuality < 0 {
			currentItem.Quality = 0
		} else {
			currentItem.Quality = tempQuality
		}
	}
	return &currentItem
}

func UpdateBackStage(currentItem Item) *Item {
	currentItem.SellIn--
	if currentItem.SellIn < 0 {
		currentItem.Quality = 0
	} else if currentItem.SellIn > 5 && currentItem.SellIn <= 10 {
		currentItem.Quality = currentItem.Quality + 2
	} else if currentItem.SellIn <= 5 && currentItem.SellIn >= 0 {
		currentItem.Quality = currentItem.Quality + 3
	} else {
		currentItem.Quality--
	}

	return &currentItem

}

func UpdateDefault(currentItem Item) *Item {
	currentItem.SellIn--

	if currentItem.SellIn < 0 {
		tempQuality := currentItem.Quality - 2
		if tempQuality > 0 {
			currentItem.Quality = tempQuality
		} else {
			currentItem.Quality = 0
		}
	} else {
		if currentItem.Quality > 0 {
			currentItem.Quality--
		}
	}
	return &currentItem
}
