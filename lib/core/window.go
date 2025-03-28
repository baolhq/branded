package core

import (
	"baolhq/branded/lib/data"
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

type Window struct {
	Content         string
	Chapter         data.Chapter
	Ready           bool
	MessageViewport viewport.Model
	ContentViewport viewport.Model
	RsoGraph        viewport.Model
	Controls        Controls
}

func InitState() Window {
	cX, cY := 12, 7 // Initial cursor position
	chapter := gen.GetChapter(cX, cY)
	gen.SeedData(&chapter)

	return Window{
		Chapter:  chapter,
		Content:  gen.RenderChapter(cX, cY, chapter),
		Controls: Controls{CursorX: cX, CursorY: cY},
	}
}

func (m Window) Init() tea.Cmd {
	return nil
}

func (w Window) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		return HandleInput(w, msg)

	case tea.WindowSizeMsg:
		if !w.Ready {
			w.ContentViewport = ui.InitContent(meta.MapWidth, meta.MapHeight, w.Content)
			w.MessageViewport = ui.InitMessage(msg.Width)
			w.RsoGraph = ui.InitRso([]int{2, 1, -2, 0, 1})
			w.Ready = true
		} else {
			w.ContentViewport.Width = meta.MapWidth
			w.ContentViewport.Height = meta.MapHeight
			w.MessageViewport.Width = meta.MsgWidth
			w.MessageViewport.Height = meta.MsgHeight
		}
	}

	return w, cmd
}

func (m Window) View() string {
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
