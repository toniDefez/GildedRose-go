package models

import "github.com/emilybache/gildedrose-refactoring-kata/gildedrose/pkg"

type ItemOld struct {
	Name            string
	SellIn, Quality int
}

type SuperItem struct {
	Name    pkg.ItemName
	SellIn  pkg.ItemSellin
	Quality pkg.ItemQuality
}

type Item interface {
	Update()
	GetName() string
	GetQuality() int
	GetSellIn() int
}

func CreateNewSuperItem(name pkg.ItemName, sellIn pkg.ItemSellin, quality pkg.ItemQuality) *SuperItem {
	return &SuperItem{
		Name:    name,
		SellIn:  sellIn,
		Quality: quality,
	}
}

func (s SuperItem) GetName() string {
	return s.Name.Name
}

func (s SuperItem) GetQuality() int {
	return s.Quality.Value
}

func (s SuperItem) GetSellIn() int {
	return s.SellIn.Value
}

func (s *SuperItem) DecreaseSellIn() {
	s.SellIn.DecreaseSellIn()
}

func (s *SuperItem) IncreaseQuality() {
	s.Quality.IncreaseQuality()
}

func (s *SuperItem) DecreaseQuality() {
	s.Quality.DecreaseQuality()
}

func (s *SuperItem) ResetQuality() {
	s.Quality.ResetQuality()
}

func (s SuperItem) IsCaducated() bool {
	return s.SellIn.IsCaducated()
}

func (s *SuperItem) DecreaseQualityByAmount(amount int) {
	s.Quality.DecreaseQualityByAmount(amount)
}

func (s SuperItem) HasToBeSoldInLessThan(days int) bool {
	return s.SellIn.IsLessThan(days)
}

// Default behaviour
func (s *SuperItem) Update() {
	s.DecreaseSellIn()
	s.DecreaseQuality()
	if s.IsCaducated() {
		s.DecreaseQuality()
	}
}
