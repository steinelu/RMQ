package main

import (
	"testing"
)

func TestSpatseTableA(t *testing.T) {
	//        				   0  1  2  3  4  5  6  7
	st := NewSparseTable([]int{1, 2, 3, 4, 5, 6, 3, 2})

	if st.Query(2, 7) != 2 {
		t.Fatalf("Error RMG(2, 7) != 2")
	}

	if st.Query(0, 7) != 1 {
		t.Fatalf("Error RMG(0, 7) != 1")
	}

	if st.Query(7, 7) != 2 {
		t.Fatalf("Error RMG(7, 7) != 2")
	}
	if st.Query(0, 0) != 1 {
		t.Fatalf("Error RMG(0, 0) != 1")
	}
}

func TestSpatseTableB(t *testing.T) {
	//        				   0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
	st := NewSparseTable([]int{1, 2, 3, 4, 5, 6, 3, 2, 9, 6, 7, 4, 7, 8, 3, 0})

	if st.Query(2, 7) != 2 {
		t.Fatalf("Error RMG(2, 7) != 2")
	}

	if st.Query(7, 15) != 0 {
		t.Fatalf("Error RMG(7, 15) != 0")
	}

	if st.Query(9, 14) != 3 {
		t.Fatalf("Error RMG(9, 14) != 3")
	}
	if st.Query(0, 15) != 0 {
		t.Fatalf("Error RMG(0, 15) != 0")
	}
}

func TestSpatseTableC(t *testing.T) {
	st := NewSparseTable([]int{})
	_, err := st.Query_(0, 0)
	if err == nil {
		t.Fatal("Error RMQ(0, 0) on an empty rmq should return error")
	}
}

func TestSpatseTableD(t *testing.T) {
	st := NewSparseTable([]int{1, 2})
	_, err := st.Query_(1, 0)
	if err == nil {
		t.Fatal("Error RMQ(1, 0) should be an error since right boarder comes before left boarder")
	}
}

func TestSpatseTableE(t *testing.T) {
	st := NewSparseTable([]int{1, 2})
	_, err := st.Query_(0, 1)
	if err != nil {
		t.Fatal("Error RMQ(0, 1) on [1, 2] should return 1")
	}
}
