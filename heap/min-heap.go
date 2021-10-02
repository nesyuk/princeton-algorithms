package main

type MinHeap []int


func NewMinHeap(array []int) *MinHeap {
	heap := MinHeap(array)
	ptr := &heap
	ptr.BuildHeap(array)
	return ptr
}

func (h MinHeap) size() int {
	return len(h) - 1
}

func (h *MinHeap) BuildHeap(array []int) {
	*h = array
	for parentIndex := h.size()/2; parentIndex >= 0; parentIndex-- {
		(*h).heapify(parentIndex)
	}
}

func (h *MinHeap) get(index int) int {
	return (*h)[index]
}

func (h *MinHeap) heapify(currentIndex int) {
	if currentIndex > h.size()/2 {
		return
	}
	minIndex := currentIndex
	leftChildIndex := currentIndex * 2 + 1
	if leftChildIndex <= h.size() && h.get(leftChildIndex) < h.get(minIndex) {
		minIndex = leftChildIndex
	}
	rigthChildIndex := leftChildIndex + 1
	if rigthChildIndex <= h.size() && h.get(rigthChildIndex) < h.get(minIndex) {
		minIndex = rigthChildIndex
	}
	if minIndex != currentIndex {
		(*h)[minIndex], (*h)[currentIndex] = (*h)[currentIndex], (*h)[minIndex]
		h.heapify(minIndex)
	}
}

func (h MinHeap) Peek() int {
	return h.get(0)
}

func (h *MinHeap) Remove() int {
	(*h)[0], (*h)[h.size()] = (*h)[h.size()], (*h)[0]
	min := (*h)[h.size()]
	(*h) = (*h)[:h.size()]
	h.heapify(0)
	return min
}

func (h *MinHeap) Insert(value int) {
	*h = append(*h, value)
	for index := h.size(); index > 0; {
		parentIndex := (index - 1)/2 
		if h.get(index) >= h.get(parentIndex) {
			break
		}
		(*h)[parentIndex], (*h)[index] = (*h)[index], (*h)[parentIndex]
		index = parentIndex
	}
}

