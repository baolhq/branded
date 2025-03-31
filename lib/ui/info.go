package ui

import (
	"baolhq/branded/lib/data"
	"baolhq/branded/lib/meta"
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

var (
	nameStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#0D92F4"))
	roleStyle      = lipgloss.NewStyle().Border(lipgloss.HiddenBorder(), false, false, true, false)
	hpStyle        = lipgloss.NewStyle().PaddingRight(1).PaddingBottom(1)
	expStyle       = lipgloss.NewStyle()
	attrStyle      = lipgloss.NewStyle().Bold(true)
	attrValueStyle = lipgloss.NewStyle().Padding(0, 1)
	itemStyle      = lipgloss.NewStyle().Border(lipgloss.ASCIIBorder(), false, false, false, true).
			PaddingLeft(1).Width(24)
	itemUsesStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#A1A1A1"))
)

// Initializes the unit info viewport
func InitInfo(u *data.Unit) viewport.Model {
	vp := viewport.New(meta.InfoWidth, meta.InfoHeight)
	vp.SetContent(renderInfo(u))
	return vp
}

// UpdateInfo unit info data
func UpdateInfo(vp *viewport.Model, u *data.Unit) {
	vp.SetContent(renderInfo(u))
}

// renderInfo returns the formatted InfoViewport
func renderInfo(u *data.Unit) string {
	var attrText, attrVal string

	name := nameStyle.Render(u.Name)
	roleText := fmt.Sprintf("%s L%d", u.Role.Name, u.Level)
	role := roleStyle.Render(roleText)
	topBlock := lipgloss.JoinVertical(lipgloss.Top, name, role)

	hpText := fmt.Sprintf("HP: %d/%d", u.Hp, u.MaxHp)
	expText := fmt.Sprintf("EXP: %d", u.Exp)
	hp := hpStyle.Render(hpText)
	exp := expStyle.Render(expText)
	hpBlock := lipgloss.JoinHorizontal(lipgloss.Left, hp, exp)

	attrText = attrStyle.Render("STR")
	attrVal = attrValueStyle.Render(strconv.Itoa(u.STR))
	strength := lipgloss.JoinHorizontal(lipgloss.Left, attrText, attrVal)

	attrText = attrStyle.Render("DEX")
	attrVal = attrValueStyle.Render(strconv.Itoa(u.DEX))
	dexterity := lipgloss.JoinHorizontal(lipgloss.Left, attrText, attrVal)

	attrText = attrStyle.Render("CON")
	attrVal = attrValueStyle.Render(strconv.Itoa(u.CON))
	constitution := lipgloss.JoinHorizontal(lipgloss.Left, attrText, attrVal)

	attrText = attrStyle.Render("INT")
	attrVal = attrValueStyle.Render(strconv.Itoa(u.INT))
	intelligence := lipgloss.JoinHorizontal(lipgloss.Left, attrText, attrVal)

	attrText = attrStyle.Render("WIS")
	attrVal = attrValueStyle.Render(strconv.Itoa(u.WIS))
	wisdom := lipgloss.JoinHorizontal(lipgloss.Left, attrText, attrVal)

	attrText = attrStyle.Render("CHA")
	attrVal = attrValueStyle.Render(strconv.Itoa(u.CHA))
	charisma := lipgloss.JoinHorizontal(lipgloss.Left, attrText, attrVal)

	leftBlock := lipgloss.JoinVertical(lipgloss.Top, strength, dexterity, constitution,
		intelligence, wisdom, charisma)

	var item, itemUses, inventory string
	itemCount := 0

	if len(u.Items) > 0 {
		itemUses = itemUsesStyle.Render(u.Items[0].UsageLeft())
		item = itemStyle.Render(u.Items[0].Name)
		inventory = lipgloss.JoinHorizontal(lipgloss.Left, item, itemUses)
		itemCount++
	}
	for i := 1; i < len(u.Items); i++ {
		itemUses = itemUsesStyle.Render(u.Items[i].UsageLeft())
		item = itemStyle.Render(u.Items[i].Name)
		temp := lipgloss.JoinHorizontal(lipgloss.Left, item, itemUses)
		inventory = lipgloss.JoinVertical(lipgloss.Top, inventory, temp)
		itemCount++
	}
	for i := itemCount + 1; i >= 0; i-- {
		item = itemStyle.Render("")
		inventory = lipgloss.JoinVertical(lipgloss.Top, inventory, item)
	}

	bottomBlock := lipgloss.JoinHorizontal(lipgloss.Left, leftBlock, inventory)
	info := lipgloss.JoinVertical(lipgloss.Top, topBlock, hpBlock, bottomBlock)

	return info
}
