package models

type Queue struct {
	TotalPages      int
	NoOfFilledPages int
	Front           *Node
	Rear            *Node
}

func CreateQueue(size int) *Queue {
	var queue Queue

	queue.TotalPages = size
	queue.NoOfFilledPages = 0
	queue.Front = nil
	queue.Rear = nil

	return &queue
}

func (q *Queue) IsQueueFull() bool {

	return q.TotalPages == q.NoOfFilledPages
}

func (q *Queue) IsQueueEmpty() bool {

	return q.NoOfFilledPages == 0 || q.Rear == nil
}

func (q *Queue) DeQueue() {
	if q.IsQueueEmpty() {
		return
	}

	if q.Front == q.Rear {
		q.Front = nil
	}

	q.Rear = q.Rear.Previous

	q.NoOfFilledPages--

}

// Enqueue adds a new pagenumber to the queue
func (q *Queue) Enqueue(h *Hash, pageNumber int) {

	// If the queue is full, remove the one at the back
	if q.IsQueueFull() {
		h.Array[q.Rear.PageNumber] = nil

		q.DeQueue()
	}

	node := NewNode(pageNumber)

	node.Next = q.Front

	if q.IsQueueEmpty() {
		q.Front = node
		q.Rear = node
	} else {
		q.Front.Previous = node
		q.Front = node
	}

	h.Array[pageNumber] = node

	q.NoOfFilledPages++

}
