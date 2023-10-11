package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	prog := tea.NewProgram(newSnakeGame(), tea.WithAltScreen())

	if err := prog.Start(); err != nil {
		panic(err)
	}
}
