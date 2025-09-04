package main

import (
	"testing"
)

func TestListManipulation(t *testing.T) {
	list = []int{}

	list = []int{}
	handleInput(5)
	expected := []int{5}
	if !equal(list, expected) {
		t.Errorf("Expected %v, got %v", expected, list)
	}

	handleInput(10)
	expected = []int{5, 10}
	if !equal(list, expected) {
		t.Errorf("Expected %v, got %v", expected, list)
	}

	handleInput(-6)
	expected = []int{9}
	if !equal(list, expected) {
		t.Errorf("Expected %v, got %v", expected, list)
	}
}

func handleInput(number int) {
	if len(list) == 0 {
		list = append(list, number)
		return
	}

	if (list[0] > 0 && number > 0) || (list[0] < 0 && number < 0) {
		list = append(list, number)
	} else {
		toReduce := abs(number)

		for toReduce > 0 && len(list) > 0 {
			if abs(list[0]) > toReduce {
				if list[0] > 0 {
					list[0] -= toReduce
				} else {
					list[0] += toReduce
				}
				toReduce = 0
			} else {
				toReduce -= abs(list[0])
				list = list[1:]
			}
		}

		if toReduce > 0 {
			list = append(list, number/abs(number)*toReduce)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
