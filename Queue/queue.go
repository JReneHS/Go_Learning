package queue

//queue para ints

import (
	"fmt"
)

type Queue struct {
	b          []string
	head, tail int
}

func (q *Queue) Push(x string) {
	switch {
	case q.tail < 0:
		next := len(q.b)
		bigger := make([]string, 2*next)
		copy(bigger[copy(bigger, q.b[q.head:]):], q.b[:q.head])
		bigger[next] = x
		q.b, q.head, q.tail = bigger, 0, next+1
	}
}

func main() {

}
