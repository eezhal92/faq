package faq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushAndPeek(t *testing.T) {
	stack := Stack{}
	q1 := Question{
		text:   "Who are you",
		answer: "I'm a bot",
	}

	q2 := Question{
		text:   "Where to apply new identity",
		answer: "Please come to gov office",
	}

	stack.Push(q1)
	stack.Push(q2)

	assert.Equal(t, stack.Peek(0).(Question), q2)
	assert.Equal(t, stack.Peek(1).(Question), q1)
}

func TestPopAndLength(t *testing.T) {
	stack := Stack{}
	if stack.Length() != 0 {
		t.Errorf("length is not right %v", stack.Length())
	}

	q1 := Question{text: "Who are you", answer: "I'm a bot"}
	q2 := Question{text: "Where to apply?", answer: "Please come to gov office"}

	stack.Push(q1)
	stack.Push(q2)

	assert.Equal(t, stack.Pop().(Question), q2)
	assert.Equal(t, stack.Length(), 1)

	stack.Push(Question{text: "a", answer: "b"})
	stack.Push(Question{text: "c", answer: "d"})

	assert.Equal(t, stack.Length(), 3)
}
