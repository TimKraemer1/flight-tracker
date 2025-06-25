package main

import (
	"fmt"
	"log"
	"time"
	"os"
	"github.com/joho/godotenv"
    "github.com/timkraemer1/flight-tracker/api"
	"github.com/timkraemer1/flight-tracker/utils"
	"github.com/timkraemer1/flight-tracker/models"
    "github.com/rivo/tview"
    "github.com/gdamore/tcell/v2"
)


func main() {
	// Create cache instance
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	cacheLocation := os.Getenv("CACHE_PATH")
	cache, err := utils.CreateSQLiteCache(cacheLocation)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	// Fetch token from .env
    token, err := api.RetrieveAuthToken()
    if err != nil {
        fmt.Printf("%v\n", err)
        return
    }
    fmt.Printf("Token Extracted\n")
    
    app := tview.NewApplication()
	pages := tview.NewPages()
	
	// Variables
	var airportCode string
	var airportName string
	var airport models.Airport
	var airportInfo string
	var arrivalInfo string
	var departureInfo string

	// Modular error window
	showModal := func(message string) {
		modal := tview.NewModal().
		SetText(message).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				pages.RemovePage("modal")
			})
		pages.AddPage("modal", modal, false, true)
	}

	// Airport info text view
	infoTextView := tview.NewTextView().
	SetDynamicColors(true).
	SetRegions(true).
	SetWordWrap(true)

	// Arrivals info
	arrivalsTextView := tview.NewTextView().
	SetDynamicColors(true).
	SetRegions(true).
	SetWordWrap(true)

	// Departures info
	departuresTextView := tview.NewTextView().
	SetDynamicColors(true).
	SetRegions(true).
	SetWordWrap(true)

	infoTextView.SetBorder(true).SetTitle("Airport Information")

	yesterday := time.Now().AddDate(0, 0, -1)
	formatted := yesterday.Format("Monday, January 02")

	arrivalsTextView.SetBorder(true).SetTitle(fmt.Sprintf("Arrivals Information (Previous Day-%s)", formatted))
	departuresTextView.SetBorder(true).SetTitle(fmt.Sprintf("Departures Information (Previous Day-%s)", formatted))

	// Text input field - first page
	inputField := tview.NewInputField().
		SetLabel("Enter Airport Code (Example, KSFO): ").
		SetFieldWidth(5).
		SetAcceptanceFunc(tview.InputFieldMaxLength(10))
	
	// First page form container
	inputForm := tview.NewForm().
		AddFormItem(inputField).
			SetFieldBackgroundColor(tcell.ColorWhite).
			SetFieldTextColor(tcell.ColorBlack).
		AddButton("Search", func() {
			airportCode = inputField.GetText()
			exists, err := false, error(nil)
			exists, airport, err = utils.AirportExists("airports.json", airportCode)
			if err != nil {
				log.Fatal(err)
			}
			if !exists {
				showModal("Airport does not exist")
			} else {
				airportName = airport.Name
				airportInfo = utils.FormatAirportInfo(airport)
				pages.SwitchToPage("list")

				// Get arrival information asynchronously
				go func() {
					arrivals, err := cache.GetArrivals(token, airportCode)
					if err != nil {
						app.QueueUpdateDraw(func() {
							showModal(fmt.Sprintf("Error: %v\n", err))
						})
						return
					}
					arrivalInfo = utils.FormatArrivals(arrivals)
					app.QueueUpdateDraw(func() {
						arrivalsTextView.SetText(arrivalInfo)
					})
				}()

				// Get departure information asynchronously
				go func() {
					departures, err := api.FetchDepartures(token, airportCode)
					if err != nil {
						app.QueueUpdateDraw(func() {
							showModal(fmt.Sprintf("Error: %v\n", err))
						})
						return
					}
					departureInfo = utils.FormatDepartures(departures)
					app.QueueUpdateDraw(func() {
						departuresTextView.SetText(departureInfo)
					})
				}()
			}
		}).
		AddButton("Quit", func() {
			app.Stop()
		}).
		SetButtonBackgroundColor(tcell.ColorWhite).
		SetButtonTextColor(tcell.ColorBlack)

	inputForm.SetBorder(true).SetTitle("Flight Tracker")

	// List menu - second page
	list := tview.NewList().
		AddItem("Departures", "List departures within 1 hour of current time", '1', nil).
		AddItem("Arrivals", "List arrivals within 1 hour of current time", '2', func() {
			pages.SwitchToPage("arrivals")
		}).
		AddItem("Airport Information", "Displays important information of airport", '3', func() {
			infoTextView.SetText(airportInfo)
			pages.SwitchToPage("information")
		}).
		AddItem("Back", "Go back to main page", 'b', func() {	
			pages.SwitchToPage("input")
		}).
		AddItem("Quit", "Exit application", 'q', func() {
			app.Stop()
		})

	list.SetBorder(true).SetTitle("Airport Options")

	// Displays airport name on top of second page
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
				textView.SetText("[yellow]Airport Code: [white]" + airportName)
			}
		}
	})

	// Add pages
	pages.AddPage("input", inputForm, true, true)
	pages.AddPage("list", listPage, true, false)
	pages.AddPage("information", infoTextView, true, false)
	pages.AddPage("arrivals", arrivalsTextView, true, false)
	pages.AddPage("departures", departuresTextView, true, false)

	// Handle global key events
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'b' {
			currentPage, _ := pages.GetFrontPage()
			if currentPage == "information" || currentPage == "arrivals" || currentPage == "departures" {
				pages.SwitchToPage("list")
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
