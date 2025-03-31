package core

import (
	"baolhq/branded/lib/ui"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// HandleInput processes player input based on state
func HandleInput(win Window, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch {
	case key.Matches(msg, Keys.Quit):
		return win, tea.Quit

	// Move cursor with direction keys
	case key.Matches(msg, Keys.Up, Keys.Down, Keys.Left, Keys.Right, Keys.UpLeft,
		Keys.UpRight, Keys.DownLeft, Keys.DownRight):
		win.Controls.MoveCursor(msg.String())

		cX, cY := win.Controls.CursorX, win.Controls.CursorY
		ui.UpdateContent(&win.ContentViewport, cX, cY, win.Chapter)

		// Update RSO graph when cursor moves onto a unit
		if unit := win.Chapter.GetUnitAt(cX, cY); unit != nil {
			ui.UpdateInfo(&win.InfoViewport, unit)
		}

	case key.Matches(msg, Keys.A):
		handleAButton(win)

	case key.Matches(msg, Keys.B):
		handleBButton(win)

	// Scroll messages
	case key.Matches(msg, Keys.ShoulderL, Keys.ShoulderR):
		ScrollMessages(&win.MessageViewport, msg)
	}

	return win, cmd
}

func handleAButton(win Window) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	sm := GetStateManager()
	state := sm.GetState()

	if state == StateIdle {
		cX, cY := win.Controls.CursorX, win.Controls.CursorY
		if unit := win.Chapter.GetUnitAt(cX, cY); unit != nil {
			sm.SetState(StateUnitSelected)
			ui.PreviewMovement(&win.ContentViewport, cX, cY, win.Chapter, *unit)
		}
	}
	return win, cmd
}

func handleBButton(win Window) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	sm := GetStateManager()
	state := sm.GetState()

	if state != StateIdle {
		sm.SetState(StateIdle)
		cX, cY := win.Controls.CursorX, win.Controls.CursorY
		ui.UpdateContent(&win.ContentViewport, cX, cY, win.Chapter)
	}
	return win, cmd
}
