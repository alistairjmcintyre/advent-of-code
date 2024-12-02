package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func combinations(slice []int) [][]int {
	combos := make([][]int, 0)
	for idx := range slice {
		combo := make([]int, len(slice))
		_ = copy(combo, slice)
		combo = remove(combo, idx)
		combos = append(combos, combo)
	}
	return combos
}

func is_sorted(slice []int) bool {
	asc := sort.IntsAreSorted(slice)
	slices.Reverse(slice)
	desc := sort.IntsAreSorted(slice)
	slices.Reverse(slice)
	return (asc || desc)
}

func pairwise(slice []int) [][]int {
	pairs := make([][]int, 0)
	prev := slice[0]
	for _, curr := range slice[1:] {
		pair := []int{prev, curr}
		pairs = append(pairs, pair)
		prev = curr
	}

	return pairs
}

func abs(a, b int) int {
	diff := a - b
	if diff < 0 {
		diff = -1 * diff
	}
	return diff
}

func is_safe(report []int) bool {
	if !is_sorted(report) {
		return false
	}

	pairs := pairwise(report)

	for _, pair := range pairs {
		difference := abs(pair[0], pair[1])
		if difference == 0 || difference > 3 {
			return false
		}
	}

	return true
}

func is_still_safe(report []int) bool {
	if is_safe(report) {
		return true
	}

	combos := combinations(report)
	for _, combo := range combos {
		if is_safe(combo) {
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open("inputs/input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	reports := make([][]int, 0)

	for scanner.Scan() {
		values := strings.Fields(scanner.Text())
		report := []int{}
		for _, element := range values {
			elem, err := strconv.Atoi(element)
			check(err)
			report = append(report, elem)
		}
		reports = append(reports, report)
	}

	safe := 0
	for _, report := range reports {
		if is_safe(report) {
			safe = safe + 1
		}
	}
	fmt.Println(safe)

	still_safe := 0
	for _, report := range reports {
		if is_still_safe(report) {
			still_safe = still_safe + 1
		}
	}

	fmt.Println(still_safe)
}
