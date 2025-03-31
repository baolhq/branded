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
			Height(meta.MapHeight).Border(lipgloss.NormalBorder(), false, true, false, false)
	infoStyle    = lipgloss.NewStyle().Width(meta.InfoWidth).Height(meta.InfoHeight)
	messageStyle = lipgloss.NewStyle().Width(meta.MsgWidth).
			Height(meta.MsgHeight).Border(lipgloss.NormalBorder(), true, false, false, false)
)

type Window struct {
	Content         string
	Chapter         data.Chapter
	Ready           bool
	MessageViewport viewport.Model
	ContentViewport viewport.Model
	RsoGraph        viewport.Model
	InfoViewport    viewport.Model
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
			w.ContentViewport = ui.InitContent(w.Content)
			w.MessageViewport = ui.InitMessage(msg.Width)

			w.InfoViewport = ui.InitInfo(&data.Unit{
				Name:  "Kros",
				Level: 20,
				Exp:   99,
				Hp:    99, MaxHp: 99,
				Role: data.Role{
					Name: "Swordmaster",
				},
				STR: 99, DEX: 99, CON: 99,
				INT: 99, WIS: 99, CHA: 99,
				Items: []*data.Item{
					{
						Name: "Bronze Sword (E)",
						Uses: 99,
					},
					{
						Name: "Iron Lance",
						Uses: 10,
					},
					{
						Name: "Vulneraries",
						Uses: -1,
					},
				},
			})

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

	panel := lipgloss.JoinVertical(lipgloss.Left,
		infoStyle.Render(m.InfoViewport.View()),
		messageStyle.Render(m.MessageViewport.View()),
	)

	content := lipgloss.JoinHorizontal(lipgloss.Top,
		mapStyle.Render(m.ContentViewport.View()),
		panel,
	)

	status := ui.RenderStatus()

	return lipgloss.JoinVertical(lipgloss.Left, content, status)
}
