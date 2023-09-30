package main

import (
	"fmt"
	"math"
	"slices"
)

type RMQ struct {
	st      SparseTable
	boarder []int
}


func NewRMQ(values []int) *RMQ {

	n := len(values)
	log_n := math.Floor(math.Log2(float64(n)))
	s := int(math.Floor(log_n / 2))

	chunks := [][]int{}
	for i := 0; i < n; i++ {
		if i%s == 0 {
			chunks = append(chunks, []int{})
		}
		chunks[len(chunks)-1] = append(chunks[len(chunks)-1], values[i])
	}
	//fmt.Println(arr)

	h := []int{}
	for _, a := range chunks {
		h = append(h, slices.Min(a))
	}
	//fmt.Println(h)

	

	return nil
}

