package main

import add_two_numbers "github.com/NGunthor/go_test/pkg/leetcode/add-two-numbers"

func main() {
	l1 := add_two_numbers.MakeList(5)
	l2 := add_two_numbers.MakeList(5)
	add_two_numbers.AddTwoNumbers(l1,l2)
}
