package table

type PollStatus int

const (
	PollWill PollStatus = iota + 1
	PollProgress
	PollDone
)

type QuestionType string

const (
	Singular    QuestionType = "S"
	Plural                   = "P"
	Description              = "D"
)

type Using int

const (
	USE Using = iota
	NotUSE
)

type View int

const (
	ViewText View = iota
	ViewImage
	ViewTextImage
)
