package ui

import (
	"baolhq/branded/lib/meta"

	"github.com/NimbleMarkets/ntcharts/canvas"
	"github.com/NimbleMarkets/ntcharts/linechart/wavelinechart"

	"github.com/charmbracelet/bubbles/viewport"
)

// Initializes the RSO viewport
func InitRso(data []int) viewport.Model {
	vp := viewport.New(meta.RsoWidth, meta.RsoHeight)
	vp.SetContent(renderRso(data))
	return vp
}

// UpdateRso RsoGraph data
func UpdateRso(vp *viewport.Model, data []int) {
	vp.SetContent(renderRso(data))
}

// renderRso returns the formatted RsoGraph
func renderRso(data []int) string {
	wlc := wavelinechart.New(meta.RsoWidth, meta.RsoHeight,
		wavelinechart.WithXYRange(1, 13, -5, 5),
		wavelinechart.WithXYSteps(1, 1))

	wlc.Clear()
	wlc.ClearAllData()

	for i := 1; i < len(data); i++ {
		wlc.Plot(canvas.Float64Point{X: float64(i), Y: float64(data[i-1])})
	}
	wlc.Plot(canvas.Float64Point{X: 2, Y: -2})

	wlc.Draw()

	return wlc.View()
}
