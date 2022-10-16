package main

import (
	"fmt"
	"strconv"
)

const fizz, buzz = "Fizz", "Buzz"
const fizzbuzz = fizz + buzz

func FizzBuzz(numbers []int) []string {
	result := make([]string, 0, len(numbers))

	for _, n := range numbers {
		switch { // if文を使って記述しても良い
		//write isFizzBuzz
		case isFizzBuzz(n):
			result = append(result, fizzbuzz)

		case isFizz(n):
			result = append(result, fizz)
		case isBuzz(n):
			result = append(result, buzz)
		default:
			result = appendNum(result, n)
		}
	}

	return result
}
func isFizzBuzz(n int) bool                { return isFizz(n) && isBuzz(n) }
func isFizz(n int) bool                    { return n%3 == 0 }
func isBuzz(n int) bool                    { return n%5 == 0 }
func appendNum(s []string, n int) []string { return append(s, strconv.Itoa(n)) }

//write isFizzBuzz

func main() {
	ret := FizzBuzz([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	fmt.Printf("%v", ret)
}
