package faq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAskShouldReturnString(t *testing.T) {
	q := Question{text: "Are you a bot?", answer: "Not really"}

	answer, _ := q.Ask("Are you a bot?").(string)

	assert.Equal(t, "Not really", answer)
}

func TestAskShouldReturnQuestionWhenHasChoices(t *testing.T) {
	q := Question{
		text:   "I'd like to have a fruit",
		answer: "Which one do you like?",
		choices: []Question{
			Question{text: "Apple", answer: "It's $4"},
			Question{text: "Grapes", answer: "It's $8"},
		},
	}

	qns, _ := q.Ask("Apple").(Question)

	assert.Equal(t, "It's $4", qns.answer)
}

func TestReplyShouldReturnReply(t *testing.T) {
	q := Question{
		text:   "I'd like to have a fruit",
		answer: "Which one do you like?",
		choices: []Question{
			Question{text: "Apple", answer: "It's $4"},
			Question{text: "Grapes", answer: "It's $8"},
		},
	}

	reply := q.Reply()

	assert.Equal(t, "Which one do you like?", reply.text)
	assert.Equal(t, 2, len(reply.choices))
}

func TestHasChoices(t *testing.T) {
	q1 := Question{text: "Hey", answer: "Ho"}

	assert.True(t, !q1.HasChoices())

	choices := make([]Question, 0)
	choices = append(choices, Question{text: "Can you help me", answer: "Probably"})

	q2 := Question{
		text:    "What are you doing?",
		answer:  "Nothing",
		choices: choices,
	}

	assert.True(t, q2.HasChoices())
}
