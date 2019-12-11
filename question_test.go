package faq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAskShouldReturnString(t *testing.T) {
	q := Question{Text: "Are you a bot?", Answer: "Not really"}

	answer, _ := q.Ask("Are you a bot?").(string)

	assert.Equal(t, "Not really", answer)
}

func TestAskShouldReturnQuestionWhenHasChoices(t *testing.T) {
	q := Question{
		Text:   "I'd like to have a fruit",
		Answer: "Which one do you like?",
		Choices: []Question{
			Question{Text: "Apple", Answer: "It's $4"},
			Question{Text: "Grapes", Answer: "It's $8"},
		},
	}

	qns, _ := q.Ask("Apple").(Question)

	assert.Equal(t, "It's $4", qns.Answer)
}

func TestReplyShouldReturnReply(t *testing.T) {
	q := Question{
		Text:   "I'd like to have a fruit",
		Answer: "Which one do you like?",
		Choices: []Question{
			Question{Text: "Apple", Answer: "It's $4"},
			Question{Text: "Grapes", Answer: "It's $8"},
		},
	}

	reply := q.Reply()

	assert.Equal(t, "Which one do you like?", reply.Text)
	assert.Equal(t, 2, len(reply.Choices))
}

func TestHasChoices(t *testing.T) {
	q1 := Question{Text: "Hey", Answer: "Ho"}

	assert.True(t, !q1.HasChoices())

	choices := make([]Question, 0)
	choices = append(choices, Question{Text: "Can you help me", Answer: "Probably"})

	q2 := Question{
		Text:    "What are you doing?",
		Answer:  "Nothing",
		Choices: choices,
	}

	assert.True(t, q2.HasChoices())
}
