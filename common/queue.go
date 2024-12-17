package common

import "log"

type Queue struct {
	Items []any
}

func (q *Queue) Enqueue(data any) {
	q.Items = append(q.Items, data)
}

func (q *Queue) Dequeue() any {
	if q.IsEmpty() {
		log.Panic("Queue is empty")
	}
	returnVal := q.Items[0]
	q.Items = q.Items[1:len(q.Items)]
	return returnVal
}

func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0
}
