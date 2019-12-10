package faq

type Driver struct {
	rootQuestion Question
	backCmd      string
	repeatCmd    string
	resetCmd     string
	stack        Stack
}

func NewDriver(rootQuestion Question, backCmd string, repeatCmd string, resetCmd string) Driver {
	d := Driver{
		rootQuestion: rootQuestion,
		backCmd:      backCmd,
		repeatCmd:    repeatCmd,
		resetCmd:     resetCmd,
		stack:        NewStack(),
	}

	d.stack.Push(rootQuestion)

	return d
}

func (d *Driver) getCommandChoices() []string {
	choices := []string{d.backCmd, d.repeatCmd, d.resetCmd}

	return choices
}

func (d *Driver) CurrentQuestion() interface{} {
	return d.stack.Peek(0)
}

func (d *Driver) PreviousQuestion() interface{} {
	return d.stack.Peek(1)
}

func (d *Driver) Boot() Reply {
	return d.rootQuestion.Reply()
}

func (d *Driver) Ask(cmd string) interface{} {
	switch cmd {
	case d.backCmd:
		return d.Back()
	case d.repeatCmd:
		return d.Repeat()
	case d.resetCmd:
		return d.Reset()
	}

	q, ok := d.CurrentQuestion().(Question)

	if !ok {
		return Reply{text: "Answer is not available"}
	}

	nextQ, _ := q.Ask(cmd).(Question)
	d.stack.Push(nextQ)

	if !nextQ.HasChoices() {
		return Reply{
			text:    nextQ.answer,
			choices: d.getCommandChoices(),
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
	d.stack.Push(d.rootQuestion)

	return d.Boot()
}
