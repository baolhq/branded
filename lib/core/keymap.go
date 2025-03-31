package core

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up, Down, Left, Right key.Binding
	A, B, X, Y            key.Binding
	Select, Start         key.Binding
	ShoulderL, ShoulderR  key.Binding
	TriggerL, TriggerR    key.Binding

	// Keyboard only
	UpLeft, UpRight     key.Binding
	DownLeft, DownRight key.Binding
	Quit                key.Binding
	q                   key.Binding
}

var Keys = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up", "8"),
		key.WithHelp("k", "D-Pad Up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down", "2"),
		key.WithHelp("j", "D-Pad Down"),
	),
	Left: key.NewBinding(
		key.WithKeys("h", "left", "4"),
		key.WithHelp("h", "D-Pad Left"),
	),
	Right: key.NewBinding(
		key.WithKeys("l", "right", "6"),
		key.WithHelp("l", "D-Pad Right"),
	),
	A: key.NewBinding(
		key.WithKeys("z"),
		key.WithHelp("z", "A Button"),
	),
	B: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "B Button"),
	),
	X: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "X Button"),
	),
	Y: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "Y Button"),
	),
	Select: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "Select Button"),
	),
	Start: key.NewBinding(
		key.WithKeys("v"),
		key.WithHelp("v", "Start Button"),
	),
	ShoulderL: key.NewBinding(
		key.WithKeys("o"),
		key.WithHelp("o", "Shoulder Left"),
	),
	ShoulderR: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "Shoulder Right"),
	),
	TriggerL: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "Trigger Left"),
	),
	TriggerR: key.NewBinding(
		key.WithKeys("w"),
		key.WithHelp("w", "Trigger Right"),
	),
	UpLeft: key.NewBinding(
		key.WithKeys("y", "7"),
		key.WithHelp("y/7", "D-Pad UpLeft"),
	),
	UpRight: key.NewBinding(
		key.WithKeys("u", "9"),
		key.WithHelp("u/9", "D-Pad UpRight"),
	),
	DownLeft: key.NewBinding(
		key.WithKeys("b", "1"),
		key.WithHelp("b/1", "D-Pad DownLeft"),
	),
	DownRight: key.NewBinding(
		key.WithKeys("n", "3"),
		key.WithHelp("n/3", "D-Pad DownRight"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "Q"),
		key.WithHelp("^C/Q", "Quit"),
	),
	q: key.NewBinding(key.WithKeys("q")),
}
