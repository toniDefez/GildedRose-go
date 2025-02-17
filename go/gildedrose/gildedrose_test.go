package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
	models "github.com/emilybache/gildedrose-refactoring-kata/gildedrose/models"
)

func Test_Foo(t *testing.T) {
	var items = []*models.ItemOld{
		{Name: "fixme", SellIn: 0, Quality: 0},
	}

	itemsNew := gildedrose.UpdateQuality(items)

	if itemsNew[0].GetName() != "fixme" {
		t.Errorf("Name: Expected %s but got %s ", "fixme", itemsNew[0].GetName())
	}
}

func Test_MaximumQuality(t *testing.T) {
	var items = []*models.ItemOld{
		{Name: "Aged Brie", SellIn: -1, Quality: 49},
	}
	itemsNew := gildedrose.UpdateQuality(items)

	if itemsNew[0].GetQuality() != 50 {
		t.Errorf("Quality : Expected %d but got %d ", 50, items[0].Quality)
	}
}

func Test_ReduceQuality(t *testing.T) {
	var items = []*models.ItemOld{
		{
			Name: "simple", SellIn: 2, Quality: 2},
	}
	itemsNew := gildedrose.UpdateQuality(items)
	if itemsNew[0].GetQuality() != 1 {
		t.Errorf("Quality : Expected %d but got %d ", 1, itemsNew[0].GetQuality())
	}
	if itemsNew[0].GetSellIn() != 1 {
		t.Errorf("SellIn : Expected %d but got %d ", 1, itemsNew[0].GetSellIn())
	}
}

func Test_Reduce_QualityDouble(t *testing.T) {
	var items = []*models.ItemOld{
		{Name: "caducated", SellIn: 0, Quality: 3},
	}

	itemsNew := gildedrose.UpdateQuality(items)

	if itemsNew[0].GetQuality() != 1 {
		t.Errorf("Quality : Expected %d but got %d ", 1, itemsNew[0].GetQuality())
	}

}

func Test_Quality_Never_Negative(t *testing.T) {
	var items = []*models.ItemOld{
		{Name: "caducated", SellIn: 0, Quality: 0},
	}

	itemsNew := gildedrose.UpdateQuality(items)

	if itemsNew[0].GetQuality() != 0 {
		t.Errorf("Quality : Expected %d but got %d ", 0, itemsNew[0].GetQuality())
	}
}

func Test_AgedBrie_IncreaseQuality(t *testing.T) {
	var items = []*models.ItemOld{
		{Name: "Aged Brie", SellIn: 2, Quality: 1},
	}

	itemsNew := gildedrose.UpdateQuality(items)

	if itemsNew[0].GetQuality() != 2 {
		t.Errorf("Quality : Expected %d but got %d ", 2, itemsNew[0].GetQuality())
	}
}

func Test_AgedBrie_IncreaseQuality2(t *testing.T) {
	var items = []*models.ItemOld{
		{Name: "Aged Brie", SellIn: 0, Quality: 1},
	}

	itemsNew := gildedrose.UpdateQuality(items)

	if itemsNew[0].GetQuality() != 3 {
		t.Errorf("Quality : Expected %d but got %d ", 3, itemsNew[0].GetQuality())
	}
}

func Test_Sulfuras_Quality(t *testing.T) {
	var items = []*models.ItemOld{
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: 25, Quality: 25},
	}

	itemsNew := gildedrose.UpdateQuality(items)

	if itemsNew[0].GetQuality() != 25 {
		t.Errorf("Quality : Expected %d but got %d ", 1, itemsNew[0].GetQuality())
	}
	if itemsNew[0].GetSellIn() != 25 {
		t.Errorf("SellIn : Expected %d but got %d ", 1, itemsNew[0].GetSellIn())
	}
}

func Test_BackStage_Ten_days(t *testing.T) {
	var items = []*models.ItemOld{
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 1},
	}

	itemsNew := gildedrose.UpdateQuality(items)

	if itemsNew[0].GetQuality() != 3 {
		t.Errorf("Quality : Expected %d but got %d ", 3, itemsNew[0].GetQuality())
	}
}

func Test_BackStage_Five_days(t *testing.T) {
	var items = []*models.ItemOld{
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 1},
	}

	itemsNew := gildedrose.UpdateQuality(items)

	if itemsNew[0].GetQuality() != 4 {
		t.Errorf("Quality : Expected %d but got %d ", 4, itemsNew[0].GetQuality())
	}
}

func Test_BackStage_Five_days_maximum(t *testing.T) {
	var items = []*models.ItemOld{
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 50},
	}

	itemsNew := gildedrose.UpdateQuality(items)

	if itemsNew[0].GetQuality() != 50 {
		t.Errorf("Quality : Expected %d but got %d ", 50, itemsNew[0].GetQuality())
	}
}

func Test_BackStage_expired(t *testing.T) {
	var items = []*models.ItemOld{
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: -1, Quality: 88},
	}

	itemsNew := gildedrose.UpdateQuality(items)

	if itemsNew[0].GetQuality() != 0 {
		t.Errorf("Quality : Expected %d but got %d ", 0, itemsNew[0].GetQuality())
	}
}

func Test_Conjured_valid(t *testing.T) {
	var items = []*models.ItemOld{
		{Name: "Conjured", SellIn: 5, Quality: 10},
	}

	itemsNew := gildedrose.UpdateQuality(items)

	if itemsNew[0].GetQuality() != 8 {
		t.Errorf("Quality : Expected %d but got %d ", 8, itemsNew[0].GetQuality())
	}
}

func Test_Conjured_Expired(t *testing.T) {
	var items = []*models.ItemOld{
		{Name: "Conjured", SellIn: 0, Quality: 10},
	}

	itemsNew := gildedrose.UpdateQuality(items)

	if itemsNew[0].GetQuality() != 6 {
		t.Errorf("Quality : Expected %d but got %d ", 6, itemsNew[0].GetQuality())
	}
}
