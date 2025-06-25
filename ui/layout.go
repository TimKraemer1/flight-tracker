package ui

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
)

type UIComponents struct {
	App						*tview.Application
	Pages					*tview.Pages
	InputForm				*tview.Form
	ArrivalsTextView		*tview.TextView
	DeparturesTextView		*tview.TextView
	InfoTextView			*tview.TextView
	ListView				*tview.List
}

func BuildUI() *UIComponents {
	return nil
}