// a circular queue
// leetcode: https://leetcode.com/problems/design-circular-queue/
package circular

type MyCircularQueue struct {
	data   []int
	head   int
	tail   int
	length int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	size := k + 1
	return MyCircularQueue{
		data:   make([]int, size, size),
		head:   0,
		tail:   0,
		length: size,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	// queue is full
	if this.IsFull() {
		return false
	}
	this.data[this.tail] = value
	this.tail++
	// hit the end
	if this.tail >= this.length {
		this.tail = 0
	}
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head++
	if this.head >= this.length {
		this.head = 0
	}
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	tail := this.tail - 1
	if tail < 0 {
		tail = this.length - 1
	}
	return this.data[tail]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == this.tail
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	tail := this.tail + 1
	if tail >= this.length {
		tail = 0
	}
	return this.head == tail
}
