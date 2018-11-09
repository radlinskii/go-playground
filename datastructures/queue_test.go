package datastructures

import "testing"

func TestMakeQueue(t *testing.T) {
	q := MakeQueue()
	if !q.list.IsEmpty() {
		t.Error("MakeQueue should create a new queue with an empty list")
	}
}

func TestEnqueue(t *testing.T) {
	q := MakeQueue()
	q.Enqueue(4)
	if q.list.GetHead().GetKey() != 4 {
		t.Error("Enqueue on an empty queue shoud add node at the front")
	}
	q.Enqueue(5)
	if q.list.GetHead().GetKey() != 4 {
		t.Error("Enqueue on non-empty queue should not modify queue's front")
	}

	if q.list.GetTail().GetKey() != 5 {
		t.Error("Enqueue should add an element at the end of the queue")
	}
}

func TestDequeue(t *testing.T) {
	q := MakeQueue()
	q.Enqueue(4)
	q.Enqueue(5)

	n := q.Dequeue()
	if n.GetKey() != 4 {
		t.Error("Dequeue should return the first value that was enqueued")
	}
	n = q.Dequeue()
	if n.GetKey() != 5 {
		t.Error("Dequeue should remove element from the queue")
	}
	n = q.Dequeue()
	if n != nil {
		t.Error("Dequeue on empty queue should return nil")
	}
}

func TestIsQueueEmpty(t *testing.T) {
	q := MakeQueue()
	if !q.IsEmpty() {
		t.Error("IsEmpty should return true on empty queue")
	}

	q.Enqueue(5)
	if q.IsEmpty() {
		t.Error("IsEmpty should return false when there are elements in the queue")
	}
}

func TestFront(t *testing.T) {
	q := MakeQueue()
	if q.Front() != nil {
		t.Error("Front on empty queue should return nil")
	}

	q.Enqueue(4)
	q.Enqueue(5)
	if q.Front().GetKey() != 4 {
		t.Error("Front should return first node from the queue")
	}

	if q.Front().GetKey() != 4 {
		t.Error("Front should not remove nodes from the queue")
	}
}

func TestQueueString(t *testing.T) {
	q := MakeQueue()
	got := q.String()
	expected := "Empty queue!"
	if got != expected {
		t.Errorf("expected: %q, got: %q", expected, got)
	}

	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(7)
	got = q.String()
	expected = "->4->5->7->"
	if got != expected {
		t.Errorf("expected: %q, got: %q", expected, got)
	}

	q.Dequeue()
	q.Dequeue()

	got = q.String()
	expected = "->7->"
	if got != expected {
		t.Errorf("expected: %q, got: %q", expected, got)
	}

}
