package main

import (
	"os"
	"strconv"

	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func main() {
	// Creating a new application
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Add dark mode
	darkMode(app)

	// Creating the main window
	window := widgets.NewQMainWindow(nil, 0)

	// Adding a menu
	addMenu(window)

	// Setting the minimum and maximum size of the window
	window.SetMinimumSize2(400, 300)
	window.SetMaximumSize2(400, 300)

	// Creating the central widget and layout
	widget := widgets.NewQWidget(nil, 0)
	window.SetCentralWidget(widget)
	layout := widgets.NewQVBoxLayout()

	// Adding selectors and labels
	for i := 1; i <= 3; i++ {
		// Creating a horizontal layout for each selector and label
		hLayout := widgets.NewQHBoxLayout()
		label := widgets.NewQLabel2("Selector "+strconv.Itoa(i)+":", nil, 0)
		comboBox := widgets.NewQComboBox(nil)
		comboBox.AddItems([]string{"Option 1", "Option 2", "Option 3"})

		// Adding the label and selector to the horizontal layout
		hLayout.AddWidget(label, 0, 0)
		hLayout.AddWidget(comboBox, 0, 0)

		// Adding the horizontal layout to the main vertical layout
		layout.AddLayout(hLayout, 0)
	}

	// Setting the main layout for the central widget
	widget.SetLayout(layout)

	// Showing the window
	window.Show()

	// Running the application
	app.Exec()

}

func darkMode(app *widgets.QApplication) {
	// Creating a dark palette
	darkPalette := gui.NewQPalette()
	darkColor := gui.NewQColor3(53, 53, 53, 255)
	textColor := gui.NewQColor3(255, 255, 255, 255)

	darkPalette.SetColor2(gui.QPalette__Window, darkColor)
	darkPalette.SetColor2(gui.QPalette__WindowText, textColor)
	darkPalette.SetColor2(gui.QPalette__Base, gui.NewQColor3(25, 25, 25, 255))
	darkPalette.SetColor2(gui.QPalette__AlternateBase, darkColor)
	darkPalette.SetColor2(gui.QPalette__ToolTipBase, textColor)
	darkPalette.SetColor2(gui.QPalette__ToolTipText, textColor)
	darkPalette.SetColor2(gui.QPalette__Text, textColor)
	darkPalette.SetColor2(gui.QPalette__Button, darkColor)
	darkPalette.SetColor2(gui.QPalette__ButtonText, textColor)
	darkPalette.SetColor2(gui.QPalette__BrightText, gui.NewQColor3(255, 0, 0, 255))
	darkPalette.SetColor2(gui.QPalette__Link, gui.NewQColor3(42, 130, 218, 255))
	darkPalette.SetColor2(gui.QPalette__Highlight, gui.NewQColor3(42, 130, 218, 255))
	darkPalette.SetColor2(gui.QPalette__HighlightedText, gui.NewQColor3(0, 0, 0, 255))

	// Applying the dark palette
	app.SetPalette(darkPalette, "")
	app.SetStyle2("Fusion")
}

func addMenu(window *widgets.QMainWindow) {
	// Creating the menu bar
	menuBar := window.MenuBar()

	// Adding the "File" menu to the menu bar
	fileMenu := menuBar.AddMenu2("File")

	// Adding actions to the "File" menu
	/*newAction := */
	fileMenu.AddAction("New")
	/*openAction := */ fileMenu.AddAction("Open")
	/*saveAction := */ fileMenu.AddAction("Save")
	fileMenu.AddSeparator()
	/*exitAction := */ fileMenu.AddAction("Exit")

	// Adding the "View" menu to the menu bar
	viewMenu := menuBar.AddMenu2("View")

	// Adding actions to the "View" menu
	statusBarAction := viewMenu.AddAction("Toggle Status Bar")

	// Example of handling the "Exit" action click event
	// exitAction.ConnectTriggered(func(checked bool) {
	//  widgets.QApplication_Quit()
	// })

	// Example of handling the "Toggle Status Bar" action click event
	statusBarAction.SetCheckable(true) // Робимо дію перевіряємою (тобто можна ввімкнути/вимкнути)
	statusBarAction.ConnectTriggered(func(checked bool) {
		// Показуємо або ховаємо статус бар залежно від стану дії
		window.StatusBar().SetVisible(checked)
	})

}
