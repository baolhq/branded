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
	enemyNameStyle = lipgloss.NewStyle().Foreground(meta.Red)
	partyNameStyle = lipgloss.NewStyle().Foreground(meta.Blue)
	allyNameStyle  = lipgloss.NewStyle().Foreground(meta.Green)
	roleStyle      = lipgloss.NewStyle().Border(lipgloss.HiddenBorder(), false, false, true, false)
	hpStyle        = lipgloss.NewStyle().PaddingRight(1).PaddingBottom(1)
	expStyle       = lipgloss.NewStyle()
	attrStyle      = lipgloss.NewStyle().Bold(true)
	attrValueStyle = lipgloss.NewStyle().Padding(0, 1)
	itemStyle      = lipgloss.NewStyle().Border(lipgloss.ASCIIBorder(), false, false, false, true).
			PaddingLeft(1).Width(meta.InfoWidth - 10)
	itemUsesStyle = lipgloss.NewStyle().Foreground(meta.Gray)
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

func renderAttribute(attr string, val int) string {
	label := attrStyle.Render(attr)
	value := attrValueStyle.Render(strconv.Itoa(val))
	return lipgloss.JoinHorizontal(lipgloss.Left, label, value)
}

func renderInventory(items []*data.Item) string {
	var item, itemUses, inventory string
	itemCount := 0

	if len(items) > 0 {
		itemUses = itemUsesStyle.Render(items[0].UsageLeft())
		item = itemStyle.Render(items[0].Name)
		inventory = lipgloss.JoinHorizontal(lipgloss.Left, item, itemUses)
		itemCount++
	}
	for i := 1; i < len(items); i++ {
		itemUses = itemUsesStyle.Render(items[i].UsageLeft())
		item = itemStyle.Render(items[i].Name)
		temp := lipgloss.JoinHorizontal(lipgloss.Left, item, itemUses)
		inventory = lipgloss.JoinVertical(lipgloss.Top, inventory, temp)
		itemCount++
	}
	for i := itemCount + 1; i >= 0; i-- {
		item = itemStyle.Render("")
		inventory = lipgloss.JoinVertical(lipgloss.Top, inventory, item)
	}

	return inventory
}

// renderInfo returns the formatted InfoViewport
func renderInfo(u *data.Unit) string {
	var name string

	switch u.Faction {
	case data.Party:
		name = partyNameStyle.Render(u.Name)
	case data.Enemy:
		name = enemyNameStyle.Render(u.Name)
	default:
		name = allyNameStyle.Render(u.Name)
	}

	roleText := fmt.Sprintf("%s L%d", u.Role.Name, u.Level)
	role := roleStyle.Render(roleText)
	topBlock := lipgloss.JoinVertical(lipgloss.Top, name, role)

	hpText := fmt.Sprintf("HP: %d/%d", u.Hp, u.MaxHp)
	expText := fmt.Sprintf("EXP: %d", u.Exp)
	hp := hpStyle.Render(hpText)
	exp := expStyle.Render(expText)
	hpBlock := lipgloss.JoinHorizontal(lipgloss.Left, hp, exp)

	str := renderAttribute("STR", u.STR)
	dex := renderAttribute("DEX", u.DEX)
	con := renderAttribute("CON", u.CON)
	intel := renderAttribute("INT", u.INT)
	wis := renderAttribute("WIS", u.WIS)
	cha := renderAttribute("STR", u.CHA)
	leftBlock := lipgloss.JoinVertical(lipgloss.Top, str, dex, con, intel, wis, cha)

	items := renderInventory(u.Items)
	bottomBlock := lipgloss.JoinHorizontal(lipgloss.Left, leftBlock, items)
	info := lipgloss.JoinVertical(lipgloss.Top, topBlock, hpBlock, bottomBlock)

	return info
}
