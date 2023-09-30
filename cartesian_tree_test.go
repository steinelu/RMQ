package main

import (
	"slices"
	"testing"
)

func TestCreateCartesianTree(t *testing.T) {
	values := []int{3, 4, 5, 1, 6, 2}
	ct := NewCartesianTree(values)

	if ct.root.value != slices.Min(values) {
		t.Fatal("Error: root element should be the same as the min of all values")
	}

	if ct.root.left.value != 3 {
		t.Fatal("left child of root should have the value 3")
	}
}

func TestEulerTour(t *testing.T) {
	values := []int{3, 4, 5, 1, 6, 2}
	ct := NewCartesianTree(values)
	tour := ct.root.eulerTour()

	t.Log(tour)

	correct := []int{1, 3, 4, 5, 4, 3, 1, 2, 6, 2, 1}
	for i := 0; i < len(tour); i++ {
		if tour[i] != correct[i] {
			t.Error("missmatch between shloud and is")
		}
	}
}
