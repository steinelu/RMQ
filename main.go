package main

func main() {
	//values := []int{1, 2, 3, 4, 5, 6, 3, 2, 1, 2, 3, 4, 5, 6, 7, 8}
	//st := NewSparseTable(values)
	values := []int{}
	for i := 0; i < 64; i++ {
		values = append(values, i)
	}
	rmq := NewRMQ(values)

	_ = rmq

	
}
