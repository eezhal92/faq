package faq

type Reply struct {
	text    string
	choices []string
}

type Question struct {
	text    string
	answer  string
	choices []Question
}

func (q *Question) getChoicesText() []string {
	texts := make([]string, 0)

	for _, choice := range q.choices {
		texts = append(texts, choice.text)
	}

	return texts
}

func (q *Question) findChoiceQuestion(text string) Question {
	var found Question

	for _, choice := range q.choices {
		if choice.text == text {
			found = choice
			break
		}
	}

	return found
}

func (q *Question) HasChoices() bool {
	return len(q.choices) > 0
}

func (q *Question) Ask(text string) interface{} {
	if !q.HasChoices() {
		return q.answer
	}

	return q.findChoiceQuestion(text)
}

func (q *Question) Reply() Reply {
	return Reply{
		text:    q.answer,
		choices: q.getChoicesText(),
	}
}
