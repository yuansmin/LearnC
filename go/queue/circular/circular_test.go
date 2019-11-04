package circular

import (
	"fmt"
	"testing"
)

// func TestCircularQueue(t *testing.T) {
// 	circularQueue := Constructor(3)
// 	if !circularQueue.EnQueue(1) {
// 		t.Errorf("enque 1 got false")
// 		return
// 	}
// 	if !circularQueue.EnQueue(2) {
// 		t.Errorf("enque 2 got false")
// 		return
// 	}
// 	if !circularQueue.EnQueue(3) {
// 		t.Errorf("enque 3 got false")
// 		return
// 	}
// 	if circularQueue.EnQueue(4) {
// 		t.Errorf("enque 4 got true")
// 		return
// 	}
// 	if circularQueue.Front() != 1 {
// 		t.Errorf("front not 1")
// 		return
// 	}
// 	if circularQueue.Rear() != 3 {
// 		t.Errorf("rear not 3")
// 		return
// 	}
// 	if !circularQueue.IsFull() {
// 		t.Errorf("not empty")
// 		return
// 	}
// 	if circularQueue.IsEmpty() {
// 		t.Errorf("not full")
// 		return
// 	}
// 	if !circularQueue.DeQueue() {
// 		t.Errorf("dequeue err")
// 		return
// 	}
// 	if !circularQueue.EnQueue(4) {
// 		t.Errorf("2 enque 4 got false")
// 		return
// 	}
// 	if circularQueue.Rear() != 4 {
// 		t.Errorf("wrong rear")
// 		return
// 	}
// 	for i := 0; i < 3; i++ {
// 		circularQueue.DeQueue()
// 	}
// 	if !circularQueue.IsEmpty() {
// 		t.Errorf("wrong empty judge")
// 	}
// }

func BenchmarkIf(b *testing.B) {
	var a int
	for i := 0; i < b.N; i++ {
		if 3500 > 3000 {
			a = 500
		}
	}
	fmt.Printf("%d\n", a)
}

func BenchmarkMod(b *testing.B) {
	var a int
	for i := 0; i < b.N; i++ {
		a = 3500 % 3000
	}
	fmt.Printf("%d\n", a)
}
