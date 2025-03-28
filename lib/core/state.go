package core

import (
	"baolhq/branded/lib/gen"
	"baolhq/branded/lib/meta"
	"baolhq/branded/lib/ui"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	mapStyle = lipgloss.NewStyle().Width(meta.MapWidth).
			Height(meta.MapHeight).Border(lipgloss.RoundedBorder())
	messageStyle = lipgloss.NewStyle().Width(meta.MsgWidth).
			Height(meta.MsgHeight).Border(lipgloss.RoundedBorder())
	rsoStyle = lipgloss.NewStyle().Width(meta.RsoWidth).
			Height(meta.RsoHeight).Border(lipgloss.HiddenBorder(), true, false, false, false)
)

type State struct {
	Content         string
	Ready           bool
	MessageViewport viewport.Model
	ContentViewport viewport.Model
	RsoGraph        viewport.Model
	Controls        Controls
}

func InitState() State {
	cX, cY := 12, 7 // Initial cursor position
	chapter := gen.GetChapter(cX, cY)

	return State{
		Content:  string(chapter),
		Controls: Controls{CursorX: cX, CursorY: cY},
	}
}

func (m State) Init() tea.Cmd {
	return nil
}

func (m State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		// Move cursor within boundaries
		case "left", "down", "up", "right",
			"h", "j", "k", "l", "y", "u", "b", "n",
			"1", "2", "3", "4", "6", "7", "8", "9":
			m.Controls.MoveCursor(msg.String(), m.ContentViewport.Width, m.ContentViewport.Height)
			newMapContent := gen.GetChapter(m.Controls.CursorX, m.Controls.CursorY)
			m.ContentViewport.SetContent(newMapContent)

		// Scroll messages
		case "o", "p", "O", "P":
			ScrollMessages(&m.MessageViewport, msg.String())

		// Update RSO on keypress
		case " ":
			ui.Update(&m.RsoGraph, []int{4, 1, -3, 2, 4})
		}

	case tea.WindowSizeMsg:
		if !m.Ready {
			m.ContentViewport = ui.InitContent(meta.MapWidth, meta.MapHeight, m.Content)
			m.MessageViewport = ui.InitMessage(msg.Width)
			m.RsoGraph = ui.InitRso([]int{2, 1, -2, 0, 1})
			m.Ready = true
		} else {
			m.ContentViewport.Width = meta.MapWidth
			m.ContentViewport.Height = meta.MapHeight
			m.MessageViewport.Width = meta.MsgWidth
			m.MessageViewport.Height = meta.MsgHeight
		}
	}

	return m, cmd
}

func (m State) View() string {
	if !m.Ready {
		return "\n  Initializing..."
	}

	rightPanel := lipgloss.JoinVertical(lipgloss.Left,
		rsoStyle.Render(m.RsoGraph.View()),
		messageStyle.Render(m.MessageViewport.View()),
	)

	topRow := lipgloss.JoinHorizontal(lipgloss.Top,
		mapStyle.Render(m.ContentViewport.View()),
		rightPanel,
	)

	bottomRow := ui.RenderStatus()

	return lipgloss.JoinVertical(lipgloss.Left, topRow, bottomRow)
}
