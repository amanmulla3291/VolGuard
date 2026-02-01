package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/amanmulla3291/volguard/internal/tui"
)

func main() {
	p := tea.NewProgram(tui.NewModel())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
