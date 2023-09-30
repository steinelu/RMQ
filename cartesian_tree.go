package main

type CartesianTreeNode struct {
	parent *CartesianTreeNode
	left   *CartesianTreeNode
	right  *CartesianTreeNode

	index int
	value int
}

type CartesianTree struct {
	root *CartesianTreeNode
}

func (self CartesianTree) eulerTour() []int {
	tour := self.root.eulerTour()
	return tour
}

func (node CartesianTreeNode) eulerTour() []int {
	tour := []int{node.value}

	if node.left != nil {
		tour = append(tour, node.left.eulerTour()...)
		tour = append(tour, node.value)
	}
	if node.right != nil {
		tour = append(tour, node.right.eulerTour()...)
		tour = append(tour, node.value)
	}
	return tour
}

func buildCartesianTree(values []int) *CartesianTreeNode {
	var cur *CartesianTreeNode = nil

	for i, v := range values {
		node := &CartesianTreeNode{nil, nil, nil, i, v}

		if cur == nil {
			cur = node
			continue
		}

		if cur.value < node.value {
			cur.right = node
			node.parent = cur
			cur = node
		} else {
			for cur.value >= node.value {
				if cur.parent == nil {
					break
				}
				cur = cur.parent
			}
			if cur.value >= node.value {
				cur.parent = node
				node.left = cur
				cur = node
			} else {
				node.parent = cur
				node.left = cur.right
				cur.right = node

				if cur.right != nil {
					node.left.parent = node
				}
				cur = node
			}
		}
	}

	for cur.parent != nil {
		cur = cur.parent
	}

	return cur
}

func NewCartesianTree(values []int) *CartesianTree {
	return &CartesianTree{buildCartesianTree(values)}
}
