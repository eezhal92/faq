package faq

type Reply struct {
	Text    string
	Choices []string
}

type Question struct {
	Text    string
	Answer  string
	Choices []Question
}

func (q *Question) getChoicesText() []string {
	texts := make([]string, 0)

	for _, choice := range q.Choices {
		texts = append(texts, choice.Text)
	}

	return texts
}

func (q *Question) findChoiceQuestion(text string) Question {
	var found Question

	for _, choice := range q.Choices {
		if choice.Text == text {
			found = choice
			break
		}
	}

	return found
}

func (q *Question) HasChoices() bool {
	return len(q.Choices) > 0
}

func (q *Question) Ask(text string) interface{} {
	if !q.HasChoices() {
		return q.Answer
	}

	return q.findChoiceQuestion(text)
}

func (q *Question) Reply() Reply {
	return Reply{
		Text:    q.Answer,
		Choices: q.getChoicesText(),
	}
}
