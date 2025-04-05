package main

import (
	"baolhq/branded/lib/util"
	"fmt"
)

func main() {
	util.EncodeData()
	units := util.DecodeUnits()
	for _, unit := range units {
		fmt.Printf("Unit ID: %d, Name: %s, Brand: %s\n", unit.Id, unit.Name, unit.Brand.Name)
	}

	// p := tea.NewProgram(
	// 	core.InitState(),
	// 	tea.WithAltScreen(),
	// 	tea.WithMouseCellMotion(),
	// )

	// if _, err := p.Run(); err != nil {
	// 	fmt.Println("Could not run program: ", err)
	// 	os.Exit(1)
	// }
}
