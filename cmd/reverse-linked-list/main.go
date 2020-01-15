package main

import (
	"fmt"
	reverse_llist "github.com/NGunthor/go_test/pkg/leetcode/reverse-llist"
)

func main() {
	head := reverse_llist.NewListNode(
		1, reverse_llist.NewListNode(2, reverse_llist.NewListNode(3, nil)))
	result := head.ReverseList()
	fmt.Println(result)
}
