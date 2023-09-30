package main

import (
	"errors"
	"math"
)

type SparseTable struct {
	table [][]int
}

func NewSparseTable(values []int) *SparseTable {
	n := float64(len(values))
	m := int(math.Ceil(math.Log2(n)))
	table := [][]int{values[:]}

	for i := 0; i < m; i++ {
		row := []int{}

		p := int(math.Pow(2, float64(i)))

		for j := 0; j+p < len(table[i]); j++ {
			a := table[i][j]
			b := table[i][j+p]
			row = append(row, min(a, b))
		}

		table = append(table, row)
	}

	st := &SparseTable{table}
	return st
}

func (st SparseTable) Query(l, r int) int {
	s := r - l + 1
	p := int(math.Floor(math.Log2(float64(s))))
	o := int(math.Pow(2, float64(p)))

	return min(st.table[p][l], st.table[p][r-o+1])
}

func (st SparseTable) Query_(l, r int) (int, error) {
	if l > r {
		return 0, errors.New("Argument Error: last argument cannot be bigger than the right argument.")
	}

	if r < 0 || l >= len(st.table[0]) {
		return 0, errors.New("Argument Error: query ranges out of bounce.")
	}
	return st.Query(l, r), nil

}
