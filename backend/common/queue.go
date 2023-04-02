package common

import (
	model "backend/model"
	"fmt"
)

type Queue struct {
	Elements []model.Image
	Size     int
}

func (q *Queue) Enqueue(item model.Image) {
	q.Elements = append(q.Elements, item)
}

func (q *Queue) Dequeue() model.Image {
	if q.IsEmpty() {
		fmt.Println("UnderFlow")
		return model.Image{}
	}
	element := q.Elements[0]
	if q.GetLength() == 1 {
		q.Elements = nil
		return element
	}
	q.Elements = q.Elements[1:]
	return element // Slice off the element once it is dequeued.
}

func (q *Queue) GetLength() int {
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *Queue) Peek() model.Image {
	if q.IsEmpty() {
		return model.Image{}
	}
	return q.Elements[0]
}
