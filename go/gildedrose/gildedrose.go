package gildedrose

import (
	models "github.com/emilybache/gildedrose-refactoring-kata/gildedrose/models"
	pkg "github.com/emilybache/gildedrose-refactoring-kata/gildedrose/pkg"
)

func UpdateQuality(items []*models.ItemOld) []models.Item {

	itemsNew := ConvertToNewSuperItems(items)

	for i := 0; i < len(itemsNew); i++ {
		itemsNew[i].Update()
	}
	return itemsNew
}

func ConvertToNewSuperItems(items []*models.ItemOld) []models.Item {
	var newItems []models.Item

	for _, oldItem := range items {
		name := pkg.ItemName{Name: oldItem.Name}
		quality := pkg.ItemQuality{Value: oldItem.Quality}
		sellIn := pkg.ItemSellin{Value: oldItem.SellIn}

		var newItem models.Item
		if name.IsAgedBrie() {
			newItem = &models.AgedBrie{SuperItem: models.SuperItem{Name: name, Quality: quality, SellIn: sellIn}}
		} else if name.IsBackstagePasses() {
			newItem = &models.Backstage{SuperItem: models.SuperItem{Name: name, Quality: quality, SellIn: sellIn}}
		} else if name.IsSulfuras() {
			newItem = &models.Sulfuras{SuperItem: models.SuperItem{Name: name, Quality: quality, SellIn: sellIn}}
		} else if name.IsConjured() {
			newItem = &models.Conjured{SuperItem: models.SuperItem{Name: name, Quality: quality, SellIn: sellIn}}
		} else {
			newItem = &models.SuperItem{Name: name, Quality: quality, SellIn: sellIn} // Default
		}

		newItems = append(newItems, newItem)
	}

	return newItems
}
