package gildedrose

type ItemOld struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*ItemOld) []Item {

	itemsNew := ConvertToNewSuperItems(items)

	for i := 0; i < len(itemsNew); i++ {
		itemsNew[i].Update()
	}
	return itemsNew
}

func ConvertToNewSuperItems(items []*ItemOld) []Item {
	var newItems []Item

	for _, oldItem := range items {
		name := ItemName{Name: oldItem.Name}
		quality := ItemQuality{Value: oldItem.Quality}
		sellIn := ItemSellin{Value: oldItem.SellIn}

		var newItem Item
		if name.IsAgedBrie() {
			newItem = &AgedBrie{SuperItem{Name: name, Quality: quality, SellIn: sellIn}}
		} else if name.IsBackstagePasses() {
			newItem = &Backstage{SuperItem{Name: name, Quality: quality, SellIn: sellIn}}
		} else if name.IsSulfuras() {
			newItem = &Sulfuras{SuperItem{Name: name, Quality: quality, SellIn: sellIn}}
		} else if name.IsConjured() {
			newItem = &Conjured{SuperItem{Name: name, Quality: quality, SellIn: sellIn}}
		} else {
			newItem = &SuperItem{Name: name, Quality: quality, SellIn: sellIn} // Default
		}

		newItems = append(newItems, newItem)
	}

	return newItems
}
