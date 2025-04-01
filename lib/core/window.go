package core

import (
	"baolhq/branded/lib/data"
	"baolhq/branded/lib/gen"
	"baolhq/branded/lib/meta"
	"baolhq/branded/lib/ui"
	"baolhq/branded/lib/util"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	mapStyle = lipgloss.NewStyle().Width(meta.MapWidth).
			Height(meta.MapHeight).Border(lipgloss.NormalBorder(), false, true, false, false)
	infoStyle    = lipgloss.NewStyle().Width(meta.InfoWidth).Height(meta.InfoHeight)
	messageStyle = lipgloss.NewStyle().Width(meta.MsgWidth).
			Height(meta.MsgHeight).Border(lipgloss.NormalBorder(), true, false, false, false)
	statusStyle = lipgloss.NewStyle().Width(meta.StatusWidth).
			Border(lipgloss.NormalBorder(), true, false, false, false)
)

type Window struct {
	Content         string
	Chapter         data.Chapter
	Ready           bool
	OldMsg          []string
	LatestMsg       string
	MessageViewport viewport.Model
	ContentViewport viewport.Model
	InfoViewport    viewport.Model
	Controls        Controls
}

func InitState() Window {
	chapter := data.Chapter{Map: data.LoadMap()}
	util.SeedData(&chapter) // NOTE: Testing purpose
	cX, cY := util.GetLordPosition(&chapter)

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
			w.ContentViewport = ui.InitContent(w.Content)
			w.LatestMsg = "Latest message"
			w.MessageViewport = ui.InitMessage(w.OldMsg, w.LatestMsg)

			lord := util.GetPartyLord(&w.Chapter)
			w.InfoViewport = ui.InitInfo(lord)

			w.Ready = true
		} else {
			w.OldMsg, w.LatestMsg = ui.UpdateMessage(w.OldMsg, w.LatestMsg)
		}
	}

	return w, cmd
}

func (m Window) View() string {
	if !m.Ready {
		return "\n  Initializing..."
	}

	panel := lipgloss.JoinVertical(lipgloss.Left,
		infoStyle.Render(m.InfoViewport.View()),
		messageStyle.Render(m.MessageViewport.View()),
	)

	content := lipgloss.JoinHorizontal(lipgloss.Top,
		mapStyle.Render(m.ContentViewport.View()),
		panel,
	)

	status := statusStyle.Render(ui.UpdateStatus())
	return lipgloss.JoinVertical(lipgloss.Left, content, status)
}

func PrintMessage(w *Window, msg string) {
	w.OldMsg, w.LatestMsg = ui.UpdateMessage(w.OldMsg, msg)
	w.MessageViewport.SetContent(w.LatestMsg)
	w.MessageViewport.GotoBottom()
}
