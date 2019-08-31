import "container/heap"

type MinHeap []int

func (h MinHeap) Less(i, j int) bool {
	return h[i] <= h[j]
}
func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Pop() interface{} {
	res := 0
	res, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return res
}
func (h *MinHeap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

type KthLargest struct {
	HeapLen int
	Heap    MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	kL := KthLargest{
		Heap:    MinHeap(nums),
		HeapLen: k,
	}
	heap.Init(&kL.Heap)
	for kL.Heap.Len() > k {
		heap.Pop(&kL.Heap)
	}
	return kL
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.Heap, val)
	if this.Heap.Len() > this.HeapLen {
		heap.Pop(&this.Heap)
	}
	return this.Heap[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
