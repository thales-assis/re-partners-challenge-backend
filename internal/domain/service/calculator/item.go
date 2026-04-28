package calculatorservice

type item struct {
	value   int // actual smallest reachable value for this residue
	residue int
}

type minHeap []item

func NewMinHeap() *minHeap {
	return &minHeap{}
}

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(item))
}

func (h *minHeap) Pop() interface{} {

	old := *h

	n := len(old)

	x := old[n-1]
	*h = old[:n-1]

	return x
}
