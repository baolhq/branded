package meta

const (
	TermWidth    = 80
	TermHeight   = 24
	StatusHeight = 2
	StatusWidth  = TermWidth
	MapWidth     = TermWidth/2 + 4
	MapHeight    = TermHeight - StatusHeight
	MsgWidth     = TermWidth/2 - 6
	MsgHeight    = MapHeight/2 - StatusHeight
	InfoWidth    = MsgWidth
	InfoHeight   = MapHeight/2 + 1
)
