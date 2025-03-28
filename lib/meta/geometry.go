package meta

const (
	TermWidth    = 80
	TermHeight   = 24
	StatusHeight = 3
	MapWidth     = TermWidth/2 + 4
	MapHeight    = TermHeight - StatusHeight - 2
	MsgWidth     = TermWidth/2 - 8
	MsgHeight    = MapHeight/2 - 2
	RsoWidth     = MsgWidth
	RsoHeight    = MapHeight/2 + 2
)
