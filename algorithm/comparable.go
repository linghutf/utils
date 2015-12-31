package algorithm

type Comparable interface {
	Len() int
	Less(i, j) bool
	Swap(i, j)
}
