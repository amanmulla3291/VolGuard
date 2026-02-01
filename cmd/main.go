package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/amanmulla3291/volguard/internal/lvm"
	"github.com/amanmulla3291/volguard/internal/tui"
)

func main() {
	provider := &lvm.MockProvider{}

	p := tea.NewProgram(tui.NewModel(provider))
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
