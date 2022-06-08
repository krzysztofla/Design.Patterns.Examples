package main

import "fmt"

type Color int

const (
	red Color = iota
	yelow
	blue
)

type Item struct {
	Name  string
	Price int
	Color Color
}

type Filter struct {
}

func (*Filter) FilterByColor(color Color, items []Item) []*Item {
	result := make([]*Item, 0)

	for i, value := range items {
		if value.Color == color {
			result = append(result, &items[i])
		}
	}
	return result
}

// to make OCP more handy we use Specification pattern

type Specification interface {
	IsSatisfied(p *Item) bool
}

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied(i *Item) bool {
	return i.Color == spec.color
}

// so now we can create filtering struct that will handle filtering in a elegant way
type BetterFiltering struct{}

func (f *BetterFiltering) Filter(
	items []Item, spec Specification) []*Item {

	result := make([]*Item, 0)

	for i, v := range items {
		if spec.IsSatisfied(&v) {
			result = append(result, &items[i])
		}
	}

	return result
}

// combine specs
type AndSpecification struct {
	first, second Specification
}

func (spec AndSpecification) IsSatisfied(p *Item) bool {
	return spec.first.IsSatisfied(p) &&
		spec.second.IsSatisfied(p)
}

func mainOC() {
	item1 := Item{"Item1", 11, red}
	item2 := Item{"Item2", 22, yelow}
	item3 := Item{"Item3", 33, yelow}

	// ---
	items := []Item{item1, item2, item3}
	f := Filter{}
	sortedItems := f.FilterByColor(blue, items)
	fmt.Println(sortedItems)

	// --- better filterin - with that interface way you can easyly follow OCP
	filterByRed := ColorSpecification{red}
	bf := BetterFiltering{}

	for _, value := range bf.Filter(items, filterByRed) {
		fmt.Println(value.Color)
	}
}
