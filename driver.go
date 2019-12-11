package faq

type Driver struct {
	RootQuestion Question
	BackCmd      string
	RepeatCmd    string
	ResetCmd     string
	stack        Stack
}

func NewDriver(rootQuestion Question, backCmd string, repeatCmd string, resetCmd string) Driver {
	d := Driver{
		RootQuestion: rootQuestion,
		BackCmd:      backCmd,
		RepeatCmd:    repeatCmd,
		ResetCmd:     resetCmd,
		stack:        NewStack(),
	}

	d.stack.Push(rootQuestion)

	return d
}

func (d *Driver) getCommandChoices() []string {
	choices := []string{d.BackCmd, d.RepeatCmd, d.ResetCmd}

	return choices
}

func (d *Driver) CurrentQuestion() interface{} {
	return d.stack.Peek(0)
}

func (d *Driver) PreviousQuestion() interface{} {
	return d.stack.Peek(1)
}

func (d *Driver) Boot() Reply {
	return d.RootQuestion.Reply()
}

func (d *Driver) Ask(cmd string) Reply {
	switch cmd {
	case d.BackCmd:
		return d.Back()
	case d.RepeatCmd:
		return d.Repeat()
	case d.ResetCmd:
		return d.Reset()
	}

	q, ok := d.CurrentQuestion().(Question)

	if !ok {
		return Reply{Text: "Answer is not available"}
	}

	nextQ, _ := q.Ask(cmd).(Question)
	d.stack.Push(nextQ)

	if !nextQ.HasChoices() {
		return Reply{
			Text:    nextQ.Answer,
			Choices: d.getCommandChoices(),
		}
	}

	return nextQ.Reply()
}

func (d *Driver) Back() Reply {
	d.stack.Pop()
	q, _ := d.CurrentQuestion().(Question)

	return q.Reply()
}

func (d *Driver) Repeat() Reply {
	q, _ := d.CurrentQuestion().(Question)

	return q.Reply()
}

func (d *Driver) Reset() Reply {
	d.stack = Stack{}
	d.stack.Push(d.RootQuestion)

	return d.Boot()
}
