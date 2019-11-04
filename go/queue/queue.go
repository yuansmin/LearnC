package queue

type RecentCounter struct {
	queue []int
	count int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: make([]int, 0, 10000),
		count: 0,
	}
}

func (this *RecentCounter) Ping(t int) int {
	for this.count > 0 && t-this.queue[0] > 3000 {
		this.queue = this.queue[1:]
		this.count--
	}
	this.queue = append(this.queue, t)
	this.count++
	return this.count
}
