package main

import (
	"fmt"
	"testing"
)

func TestAppendCase1(t *testing.T) {
	fmt.Println("\nCase 1. Add 1 el")
	testArray := [1]int{}

	s := Init(50, 64)
	sOriginal := make([]int, 50, 64)

	printMyData(s, sOriginal)

	s = Append(s, testArray[:]...)
	sOriginal = append(sOriginal, testArray[:]...)

	printMyResults(s, sOriginal)
}

func TestAppendCase2(t *testing.T) {
	fmt.Println("\nCase 2. Add 15 el")
	testArray := [15]int{}

	s := Init(50, 64)
	sOriginal := make([]int, 50, 64)

	printMyData(s, sOriginal)

	s = Append(s, testArray[:]...)
	sOriginal = append(sOriginal, testArray[:]...)

	printMyResults(s, sOriginal)
}

func TestAppendCase3(t *testing.T) {
	fmt.Println("\nCase 3. Add 80 el")
	testArray := [80]int{}

	s := Init(50, 64)
	sOriginal := make([]int, 50, 64)

	printMyData(s, sOriginal)

	s = Append(s, testArray[:]...)
	sOriginal = append(sOriginal, testArray[:]...)

	printMyResults(s, sOriginal)
}

func TestAppendCase4(t *testing.T) {
	fmt.Println("\nCase 4. Add 1 el but old cap == old len")
	testArray := [1]int{}

	s := Init(50, 50)
	sOriginal := make([]int, 50, 50)

	printMyData(s, sOriginal)

	s = Append(s, testArray[:]...)
	sOriginal = append(sOriginal, testArray[:]...)

	printMyResults(s, sOriginal)
}

func TestAppendCase5(t *testing.T) {
	fmt.Println("\nCase 5. Add 15 el but old cap == old len")
	testArray := [15]int{}

	s := Init(50, 50)
	sOriginal := make([]int, 50, 50)

	printMyData(s, sOriginal)

	s = Append(s, testArray[:]...)
	sOriginal = append(sOriginal, testArray[:]...)

	printMyResults(s, sOriginal)
}

func TestAppendCase6(t *testing.T) {
	fmt.Println("\nCase 6. Add 80 el but old cap == old len")
	testArray := [80]int{}

	s := Init(50, 50)
	sOriginal := make([]int, 50, 50)

	printMyData(s, sOriginal)

	s = Append(s, testArray[:]...)
	sOriginal = append(sOriginal, testArray[:]...)

	printMyResults(s, sOriginal)
}

func printMyResults(s *Slice, sOriginal []int) {
	fmt.Println("\nMy Slice after Append:     ", s)
	fmt.Println("Original Slice after append: ", sOriginal, len(sOriginal), cap(sOriginal))
}

func printMyData(s *Slice, sOriginal []int) {
	fmt.Println("\nMy Slice:     ", s)
	fmt.Println("Original Slice: ", sOriginal, len(sOriginal), cap(sOriginal))
}
