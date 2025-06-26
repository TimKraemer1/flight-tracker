package main

import (
	"fmt"
	"time"
	"os"
	"github.com/joho/godotenv"
    "github.com/timkraemer1/flight-tracker/api"
	"github.com/timkraemer1/flight-tracker/utils"
	"github.com/timkraemer1/flight-tracker/models"
	"github.com/timkraemer1/flight-tracker/ui"
)

type AppState struct {
	token 				string
	cache 				*utils.SQLiteFlightCache
	components			*ui.UIComponents
	airport				models.Airport
	airportInfo			string
	backStack			[]string
}

func handleSearch(state *AppState) error {
	airportCode := state.components.InputField.GetText()

	// Check if airport code exists
	exists, airport, err := utils.AirportExists("./airports.json", airportCode)
	if err != nil {
		return err
	}

	if !exists {
		ui.ShowModal(state.components.Pages, "Airport does not exists (Did you remember to capitalize?)")
		return nil
	}

	// Store airport information
	state.airport = airport
	state.airportInfo = utils.FormatAirportInfo(state.airport)
	state.components.Pages.SwitchToPage("list")

	// Run asynchronously fetching the arrivals and departures
	go loadArrivals(state, airportCode)
	go loadDepartures(state, airportCode)
	return nil
}

func loadArrivals(state *AppState, airportCode string){
	arrivals, err := state.cache.GetArrivals(state.token, state.airport.Icao)
	if err != nil {
		state.components.App.QueueUpdateDraw(func() {
			ui.ShowModal(state.components.Pages, fmt.Sprintf("Error: %v\n", err))
		})
		return
	}
	arrivalInfo := utils.FormatArrivals(arrivals)
	state.components.App.QueueUpdateDraw(func() {
		state.components.ArrivalsTextView.SetText(arrivalInfo)
	})
}

func loadDepartures(state *AppState, airportCode string){
	departures, err := state.cache.GetDepartures(state.token, state.airport.Icao)
	if err != nil {
		state.components.App.QueueUpdateDraw(func() {
			ui.ShowModal(state.components.Pages, fmt.Sprintf("Error: %v\n", err))
		})
		return
	}
	departureInfo := utils.FormatDepartures(departures)
	state.components.App.QueueUpdateDraw(func() {
		state.components.DeparturesTextView.SetText(departureInfo)
	})
}

func setupDependencies() (string, *utils.SQLiteFlightCache, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", nil, err
	}
	
	token, err := api.RetrieveAuthToken()
	if err != nil {
		return "", nil, err
	}

	cache, err := utils.CreateSQLiteCache(os.Getenv("CACHE_PATH"))
	if err != nil {
		return "", nil, err
	}
	return token, cache, nil
}

func setupPages(state *AppState) {
	pages := state.components.Pages

	// Create list page
	listPage, textView, list := ui.CreateListPageLayout(state.airport.Name)
	ui.SetMenuHandlers(list, pages, state.components.App, state.components.InfoTextView, &state.airportInfo)

	// Setup page change handler
	pages.SetChangedFunc(func() {
		if pages.HasPage("list") {
			currentPage, _ := pages.GetFrontPage()
			if currentPage == "list" {
				textView.SetText("[yellow]Airport: [white]" + state.airport.Name)
			}
		}
	})

	// Add all pages
	pages.AddPage("list", listPage, true, false)
	pages.AddPage("input", state.components.InputForm, true, true)
	pages.AddPage("information", state.components.InfoTextView, true, false)
	pages.AddPage("arrivals", state.components.ArrivalsTextView, true, false)
	pages.AddPage("departures", state.components.DeparturesTextView, true, false)
}

func setupTextView(components *ui.UIComponents) {
	yesterday := time.Now().AddDate(0, 0, -1)
	formatted := yesterday.Format("Monday, January 02")

	components.InfoTextView.SetBorder(true).SetTitle("Airport Information")
	components.ArrivalsTextView.SetTitle(fmt.Sprintf("Arrivals Information (Previous Day-%s)", formatted))
	components.DeparturesTextView.SetTitle(fmt.Sprintf("Departures Information (Previous Day-%s)", formatted))
}

func setupInputForm(state *AppState) {
	searchHandler := func() {
		err := handleSearch(state)
		if err != nil {
			ui.ShowModal(state.components.Pages, fmt.Sprintf("Error: %v", err))
		}
	}

	quitHandler := func() {
		state.components.App.Stop()
	}

	state.components.SetInputFormatHandlers(searchHandler, quitHandler)
}

func main() {
	// Get token and cache path
	token, cache, err := setupDependencies()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}


	components := ui.BuildUI()
	app := components.App
	pages := components.Pages

	state := &AppState{
		token: token,
		cache: cache,
		components: components,
	}

	setupInputForm(state)
	setupPages(state)
	setupTextView(components)
	
	ui.SetupGlobalKeyHandler(app, pages)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}