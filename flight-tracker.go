package main

import (
	"fmt"
    "github.com/timkraemer1/flight-tracker/api"
    // "github.com/timkraemer1/flight-tracker/models"
    "github.com/rivo/tview"
    "github.com/gdamore/tcell/v2"
)


func main() {
    _, err := api.RetrieveAuthToken()
    if err != nil {
        fmt.Printf("%v\n", err)
        return
    }
    fmt.Printf("Token Extracted\n")
    
    app := tview.NewApplication()
	pages := tview.NewPages()
	
	// Variable to store the text from the input field
	var storedText string

	// Create the first page with text input
	inputField := tview.NewInputField().
		SetLabel("Enter Airport Code: ").
		SetFieldWidth(5).
		SetAcceptanceFunc(tview.InputFieldMaxLength(50))

	inputForm := tview.NewForm().
		AddFormItem(inputField).
		AddButton("Search", func() {
			storedText = inputField.GetText()
			pages.SwitchToPage("list")
		}).
		AddButton("Quit", func() {
			app.Stop()
		})

	inputForm.SetBorder(true).SetTitle("Airport Information")

	// Create the second page with list
	list := tview.NewList().
		AddItem("Departures", "List departures within 1 hour of current time", '1', nil).
		AddItem("Arrivals", "List arrivals within 1 hour of current time", '2', nil).
		AddItem("Airport Information", "", '3', nil).
		AddItem("Back", "Go back to main page", 'b', func() {
			pages.SwitchToPage("input")
		}).
		AddItem("Quit", "Exit application", 'q', func() {
			app.Stop()
		})

	list.SetBorder(true).SetTitle("Airport Information")

	// Add a text view to display the stored text
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	// Create a flex layout for the list page
	listPage := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(textView, 1, 0, false).
		AddItem(list, 0, 1, true)

	// Update the text view when switching to list page
	pages.SetChangedFunc(func() {
		if pages.HasPage("list") {
			currentPage, _ := pages.GetFrontPage()
			if currentPage == "list" {
				if storedText != "" {
					textView.SetText("[yellow]Airport Code: [white]" + storedText)
				} else {
					textView.SetText("[red]No text entered")
				}
			}
		}
	})

	// Add pages
	pages.AddPage("input", inputForm, true, true)
	pages.AddPage("list", listPage, true, false)

	// Handle global key events
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			currentPage, _ := pages.GetFrontPage()
			if currentPage == "list" {
				pages.SwitchToPage("input")
				return nil
			}
		}
		return event
	})

	// Run the application
	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
