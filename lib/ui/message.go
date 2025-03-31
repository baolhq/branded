package ui

import (
	"baolhq/branded/lib/meta"
	"fmt"
	"regexp"
	"strconv"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

var (
	baseMsgStyle    = lipgloss.NewStyle().Width(meta.MsgWidth)
	oldMessageStyle = baseMsgStyle.Foreground(meta.Gray)
	newMessageStyle = baseMsgStyle.Foreground(meta.White)
)

// InitMessage initializes the message viewport
func InitMessage(oldMsg []string, latestMsg string) viewport.Model {
	vp := viewport.New(meta.MsgWidth, meta.MsgHeight)
	_, content := UpdateMessage(oldMsg, latestMsg)

	vp.SetContent(content)
	vp.GotoBottom()
	return vp
}

// UpdateMessage adds a new message while keeping history manageable
func UpdateMessage(oldMsg []string, newMsg string) ([]string, string) {
	if len(oldMsg) == 0 {
		return []string{newMsg}, newMessageStyle.Render(newMsg)
	}

	lastIndex := len(oldMsg) - 1
	lastMsg := oldMsg[lastIndex]

	// Check if the new message is the same as the last one
	var count int = 2
	baseMsg := lastMsg
	match := regexp.MustCompile(`^(.*) \(x(\d+)\)$`).FindStringSubmatch(lastMsg)
	if match != nil {
		baseMsg = match[1]
		if parsedCount, err := strconv.Atoi(match[2]); err == nil {
			count = parsedCount + 1
		}
	}

	if newMsg == baseMsg {
		// Replace last message with updated count
		oldMsg[lastIndex] = fmt.Sprintf("%s (x%d)", baseMsg, count)
	} else {
		oldMsg = append(oldMsg, newMsg)
	}

	if len(oldMsg) > meta.MaxMsgHistory {
		oldMsg = oldMsg[len(oldMsg)-meta.MaxMsgHistory:] // Trim oldest messages
	}

	// Render old messages
	var renderedMsgs []string
	for _, msg := range oldMsg[:len(oldMsg)-1] { // Render all *except* the latest
		renderedMsgs = append(renderedMsgs, oldMessageStyle.Render(msg))
	}

	// Append the newest message at the bottom
	renderedMsgs = append(renderedMsgs, newMessageStyle.Render(oldMsg[len(oldMsg)-1]))

	res := lipgloss.JoinVertical(lipgloss.Top, renderedMsgs...)
	return oldMsg, res
}
