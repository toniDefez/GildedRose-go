package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Foo(t *testing.T) {
	var items = []*gildedrose.Item{
		{"fixme", 0, 0},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Name != "fixme" {
		t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].Name)
	}
}

func Test_MaximumQuality(t *testing.T) {
	var items = []*gildedrose.Item{
		{
			"Aged Brie", -1, 49},
	}
	gildedrose.UpdateQuality(items)
	if items[0].Quality != 50 {
		t.Errorf("Quality : Expected %d but got %d ", 50, items[0].Quality)
	}
}

func Test_ReduceQuality(t *testing.T) {
	var items = []*gildedrose.Item{
		{
			"simple", 2, 2},
	}
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 1 {
		t.Errorf("Quality : Expected %d but got %d ", 1, items[0].Quality)
	}
	if items[0].SellIn != 1 {
		t.Errorf("SellIn : Expected %d but got %d ", 1, items[0].SellIn)
	}
}

func Test_Reduce_QualityDouble(t *testing.T) {
	var items = []*gildedrose.Item{
		{"caducated", 0, 3},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Quality != 1 {
		t.Errorf("Quality : Expected %d but got %d ", 1, items[0].Quality)
	}

}

func Test_Quality_Never_Negative(t *testing.T) {
	var items = []*gildedrose.Item{
		{"caducated", 0, 0},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Quality != 0 {
		t.Errorf("Quality : Expected %d but got %d ", 0, items[0].Quality)
	}
}

func Test_AgedBrie_IncreaseQuality(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", 2, 1},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Quality != 2 {
		t.Errorf("Quality : Expected %d but got %d ", 2, items[0].Quality)
	}

}

func Test_AgedBrie_IncreaseQuality2(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", 0, 1},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Quality != 3 {
		t.Errorf("Quality : Expected %d but got %d ", 3, items[0].Quality)
	}
}

func Test_Sulfuras_Quality(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 25, 25},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Quality != 25 {
		t.Errorf("Quality : Expected %d but got %d ", 1, items[0].Quality)
	}
	if items[0].SellIn != 25 {
		t.Errorf("SellIn : Expected %d but got %d ", 1, items[0].SellIn)
	}
}

func Test_BackStage_Ten_days(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 10, 1},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Quality != 3 {
		t.Errorf("Quality : Expected %d but got %d ", 3, items[0].Quality)
	}
}

func Test_BackStage_Five_days(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 5, 1},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Quality != 4 {
		t.Errorf("Quality : Expected %d but got %d ", 4, items[0].Quality)
	}
}

func Test_BackStage_Five_days_maximum(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 5, 50},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Quality != 50 {
		t.Errorf("Quality : Expected %d but got %d ", 50, items[0].Quality)
	}
}

func Test_BackStage_expired(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", -1, 88},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Quality != 0 {
		t.Errorf("Quality : Expected %d but got %d ", 0, items[0].Quality)
	}
}

func Test_Conjured_valid(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Conjured", 5, 10},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Quality != 8 {
		t.Errorf("Quality : Expected %d but got %d ", 8, items[0].Quality)
	}
}

func Test_Conjured_Expired(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Conjured", 0, 10},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Quality != 6 {
		t.Errorf("Quality : Expected %d but got %d ", 6, items[0].Quality)
	}
}
