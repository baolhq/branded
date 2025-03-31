package core

import (
	"baolhq/branded/lib/ui"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// HandleInput processes player input based on state
func HandleInput(w Window, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch {
	case key.Matches(msg, Keys.q):
		PrintMessage(&w, "Press Q to quit")

	case key.Matches(msg, Keys.Quit):
		return w, tea.Quit

	// Move cursor with direction keys
	case key.Matches(msg, Keys.Up, Keys.Down, Keys.Left, Keys.Right,
		Keys.UpLeft, Keys.UpRight, Keys.DownLeft, Keys.DownRight):
		w.Controls.MoveCursor(msg.String())

		cX, cY := w.Controls.CursorX, w.Controls.CursorY
		ui.UpdateContent(&w.ContentViewport, cX, cY, w.Chapter)

		// Update RSO graph when cursor moves onto a unit
		if unit := w.Chapter.GetUnitAt(cX, cY); unit != nil {
			ui.UpdateInfo(&w.InfoViewport, unit)
		}

	case key.Matches(msg, Keys.A):
		handleAButton(&w)

	case key.Matches(msg, Keys.B):
		handleBButton(&w)

	case key.Matches(msg, Keys.ShoulderL):
		handleShoulderL(&w)

	case key.Matches(msg, Keys.ShoulderR):
		handleShoulderR(&w)
	}

	return w, cmd
}

func handleAButton(w *Window) (tea.Model, tea.Cmd) {
	sm := GetStateManager()
	state := sm.GetState()

	if state == StateIdle {
		cX, cY := w.Controls.CursorX, w.Controls.CursorY
		if unit := w.Chapter.GetUnitAt(cX, cY); unit != nil {
			sm.SetState(StateUnitSelected)
			ui.PreviewMovement(&w.ContentViewport, cX, cY, w.Chapter, *unit)
			PrintMessage(w, "Unit selected")
		}
	}
	return w, nil
}

func handleBButton(w *Window) (tea.Model, tea.Cmd) {
	sm := GetStateManager()
	state := sm.GetState()

	if state != StateIdle {
		sm.SetState(StateIdle)
		cX, cY := w.Controls.CursorX, w.Controls.CursorY
		ui.UpdateContent(&w.ContentViewport, cX, cY, w.Chapter)
		PrintMessage(w, "Unit deselected")
	}
	return w, nil
}

func handleShoulderL(w *Window) (tea.Model, tea.Cmd) {
	maxYOffset := w.MessageViewport.TotalLineCount() - w.MessageViewport.Height
	if w.MessageViewport.YOffset == 0 {
		w.MessageViewport.SetYOffset(maxYOffset)
	} else {
		w.MessageViewport.LineUp(1)
	}
	return w, nil
}

func handleShoulderR(w *Window) (tea.Model, tea.Cmd) {
	maxYOffset := w.MessageViewport.TotalLineCount() - w.MessageViewport.Height
	if w.MessageViewport.YOffset >= maxYOffset {
		w.MessageViewport.SetYOffset(0)
	} else {
		w.MessageViewport.LineDown(1)
	}
	return w, nil
}
