package gildedrose_test

import (
	"testing"

	gildedrose "github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Foo(t *testing.T) {
	var items = []*gildedrose.ItemOld{
		{"fixme", 0, 0},
	}

	itemsNew := gildedrose.UpdateQuality(items)

	if itemsNew[0].GetName() != "fixme" {
		t.Errorf("Name: Expected %s but got %s ", "fixme", itemsNew[0].GetName())
	}
}
