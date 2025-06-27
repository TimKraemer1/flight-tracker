package ui

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
)

// Modular error window
func ShowModal(pages *tview.Pages, message string) {
	modal := tview.NewModal().
		SetText(message).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.RemovePage("modal")
		})
	pages.AddPage("modal", modal, false, true)
}

// Create the main menu list
func CreateMenuList() *tview.List {
	return tview.NewList().
		AddItem("Airport Information", "Displays important information of airport", '1', nil).
		AddItem("Departures", "List departures for past day", '2', nil).
		AddItem("Arrivals", "List arrivals for past day", '3', nil).
		AddItem("Menu", "Go back to main menu", 'm', nil).
		AddItem("Quit", "Exit application", 'q', nil)
}

// Set the handlers for the menu list items
func SetMenuHandlers(list *tview.List, pages *tview.Pages, app *tview.Application, infoTextView *tview.TextView, airportInfo *string) {
	list.SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		switch index {
		case 0:
			infoTextView.SetText(*airportInfo)
			pages.SwitchToPage("information")
		case 1:
			pages.SwitchToPage("departures")
		case 2:
			pages.SwitchToPage("arrivals")
		case 3:
			pages.SwitchToPage("input")
		case 4:
			app.Stop()
		}
	})
}

// Set up global key even handling
func SetupGlobalKeyHandler(app *tview.Application, pages *tview.Pages) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'b' {
			currentPage, _ := pages.GetFrontPage()
			switch currentPage {
			case "information", "arrivals", "departures":
				pages.SwitchToPage("list")
			}
		}
		return event
	})
}

// Create the layout for the list page
func CreateListPageLayout(airportName string) (*tview.Flex, *tview.TextView, *tview.List) {
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	list := CreateMenuList()
	list.SetBorder(true).SetTitle("Airport Options")

	listPage := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(textView, 1, 0, false).
		AddItem(list, 0, 1, true)
	
	return listPage, textView, list
}