package main

import (
  "bufio"
  "fmt"
  //"io"
  "os"
  "strings"
  "strconv"
  "sort"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}


func p1(left []int, right []int) {
  // Computes the distance between each value in 2 lists of integers.
  total_dist := 0

  for idx := range len(left) {
    dist := left[idx] - right[idx]
    if dist < 0 {
      dist = -1 * dist 
    }
    total_dist += dist
  }

  fmt.Println(total_dist)
}


func p2(left []int, right []int) {
  // Calculates the "similarity" of two lists by multiplying each value in the left array
  // by the total number of occurrences of that value in the right array.
  total_similarity := 0
  
  // Store the counts in a map
  counts := make(map[int] int)
  for _ , num := range right {
    counts[num] = counts[num] + 1
  }

  for _, num := range left {
    score := num * counts[num] 
    total_similarity += score
  }

  fmt.Println(total_similarity)
  
}


func main() {
  file, err := os.Open("inputs/input.txt")
  check(err)
  defer file.Close()

  scanner := bufio.NewScanner(file)

  left  := []int{}
  right := []int{}

  for scanner.Scan() {
    values := strings.Fields(scanner.Text())

    num1, err1 := strconv.Atoi(values[0])
    check(err1)

    num2, err2 := strconv.Atoi(values[1])
    check(err2)

    left = append(left, num1)
    right = append(right, num2)
  }

  sort.Slice(left, func(i, j int) bool {
    return left[i] > left[j]
  })

  sort.Slice(right, func(i, j int) bool {
    return right[i] > right[j]
  })

  p1(left, right)
  p2(left, right)
}
