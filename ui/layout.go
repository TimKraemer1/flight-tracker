package ui

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
)

type UIComponents struct {
	App						*tview.Application
	Pages					*tview.Pages
	InputField				*tview.InputField
	InputForm				*tview.Form
	InputFormPage			*tview.Flex
	ArrivalsTextView		*tview.TextView
	DeparturesTextView		*tview.TextView
	InfoTextView			*tview.TextView
	ListView				*tview.List
}

func BuildUI() *UIComponents {
	// Create the tview application and pages
	app := tview.NewApplication()
	pages := tview.NewPages()

	// infoText view, for airport information
	infoTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	// arrivalsText view, for airport arrivals
	arrivalsTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	// departuresText view, for airport departures
	departuresTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	// inputField for inputForm
	inputField := tview.NewInputField().
		SetLabel("Enter Airport Code (Example, KSFO): ").
		SetFieldWidth(5).
		SetAcceptanceFunc(tview.InputFieldMaxLength(10))

	inputForm := createInputForm(inputField)
	 
	mainMenuHeader := tview.NewTextView().
		SetText("\n[::b]Flight Tracker Console[::-]\n[gray]Enter a valid ICAO airport code below to retrieve real-time arrival and departure data.").
		SetTextColor(tcell.ColorLightSkyBlue).
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)
	
	inputFormPage := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(mainMenuHeader, 4, 0, false).
		AddItem(inputForm, 0, 1, true)
		
	return &UIComponents{
		App: app,
		Pages: pages,
		InputField: inputField,
		InputForm: inputForm,
		InputFormPage: inputFormPage,
		ArrivalsTextView: arrivalsTextView,
		DeparturesTextView: departuresTextView,
		InfoTextView: infoTextView,
		ListView: nil,
	}
}

// Create and configure the input form
func createInputForm(inputField *tview.InputField) *tview.Form {
	form := tview.NewForm().
		AddFormItem(inputField).
		SetFieldBackgroundColor(tcell.ColorWhite).
		SetFieldTextColor(tcell.ColorBlack).
		SetButtonBackgroundColor(tcell.ColorWhite).
		SetButtonTextColor(tcell.ColorBlack)

	form.SetBorder(true)
	form.SetBorderColor(tcell.ColorGrey)

	return form
}

// Set the button handlers for the input form
func (ui *UIComponents) SetInputFormatHandlers(searchHandler, quitHandler func()) {
	ui.InputForm.AddButton("Search", searchHandler)
	ui.InputForm.AddButton("Quit", quitHandler)
}
