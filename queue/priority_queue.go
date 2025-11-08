package queue

import "errors"

type PriorityQueue struct {
	// Slice of ArrayQueues, where each index represents a priority level
	// Higher index = higher priority
	queues []ArrayQueue
	size   int
}

func NewPriorityQueue(priorityLevels int) *PriorityQueue {
	// Initialize with specified number of priority levels
	// Each queue can hold up to 100 items (configurable)
	queues := make([]ArrayQueue, priorityLevels)
	for i := 0; i < priorityLevels; i++ {
		queues[i] = New(100) // Create ArrayQueue with capacity 100
	}
	
	return &PriorityQueue{
		queues: queues,
		size:   0,
	}
}

func (pq *PriorityQueue) Clear() {
	for i := range pq.queues {
		pq.queues[i].Clear()
	}
	pq.size = 0
}

func (pq *PriorityQueue) Size() int {
	return pq.size
}


func (pq *PriorityQueue) Dequeue() (int, error) {
	if pq.size == 0 {
		return 0, ErrQueueEmpty
	}
	
	// Dequeue from the highest priority non-empty queue
	for i := len(pq.queues) - 1; i >= 0; i-- {
		if pq.queues[i].Size() > 0 {
			value, err := pq.queues[i].Dequeue()
			if err == nil {
				pq.size--
			}
			return value, err
		}
	}
	
	// This should never happen if size > 0
	return 0, ErrQueueEmpty
}

func (pq *PriorityQueue) Front() (int, error) {
	if pq.size == 0 {
		return 0, ErrQueueEmpty
	}
	
	// Get front element from the highest priority non-empty queue
	for i := len(pq.queues) - 1; i >= 0; i-- {
		if pq.queues[i].Size() > 0 {
			return pq.queues[i].Front()
		}
	}
	
	// This should never happen if size > 0
	return 0, ErrQueueEmpty
}

// Add with specific priority
func (pq *PriorityQueue) EnqueueWithPriority(e int, priority int) error {
	// Validate priority level
	if priority < 0 || priority >= len(pq.queues) {
		return errors.New("invalid priority level")
	}
	
	// Add the element to the queue with specified priority
	err := pq.queues[priority].Enqueue(e)
	if err == nil {
		pq.size++
	}
	return err
}

// Implementation of the Queue interface
func (pq *PriorityQueue) Enqueue(e int) error {
	// Default to middle priority level
	defaultPriority := len(pq.queues) / 2
	return pq.EnqueueWithPriority(e, defaultPriority)
}
