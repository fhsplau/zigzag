package main

import "github.com/szak18/zigzag/heap"
import "fmt"

func main() {
	arr := []int{10, 2, 8, 12, 14, 1, 5, -2, 7, 15}
	heap := heap.Build(arr, heap.MinHeap{})

	fmt.Println(heap)
}
