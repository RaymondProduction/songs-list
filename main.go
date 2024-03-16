package main

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	_ "github.com/glebarez/go-sqlite"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var currentDatabasePath = "./songs.db"

func main() {

	gtk.Init(nil)
	win := initGTKWindow()

	setColorWindow(win)

	win.ShowAll()

	gtk.Main()

}

func initGTKWindow() *gtk.Window {

	// Create builder
	builder, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Error bulder:", err)
	}

	populateSongsAsync(builder)

	// Lload the window from the Glade file into the builder
	err = builder.AddFromFile("main.glade")
	if err != nil {
		log.Fatal("Error when loading glade file:", err)
	}

	// We get the object of the main window by ID
	obj, err := builder.GetObject("main-window")
	if err != nil {
		log.Fatal("Error:", err)
	}

	win, ok := obj.(*gtk.Window)
	if !ok {
		log.Fatal("Failed to get window")
	}

	obj, err = builder.GetObject("open-menu-item")
	if err != nil {
		log.Fatal("Could not find menu item 'Open':", err)
	}

	openMenuItem, ok := obj.(*gtk.MenuItem)
	if !ok {
		log.Fatal("Could not get menu item 'Open'")
	}

	openMenuItem.Connect("activate", func() {
		// dialog, err := gtk.FileChooserDialogNewWith2Buttons("Виберіть файл...", win, gtk.FILE_CHOOSER_ACTION_OPEN, "Скасувати", gtk.RESPONSE_CANCEL, "Відкрити", gtk.RESPONSE_ACCEPT)
		dialog, err := gtk.FileChooserDialogNewWith2Buttons("Select file", win, gtk.FILE_CHOOSER_ACTION_OPEN, "Cancel", gtk.RESPONSE_CANCEL, "Open", gtk.RESPONSE_ACCEPT)
		if err != nil {
			log.Fatal("Failed to create file open dialog:", err)
		}
		defer dialog.Destroy()

		filter, err := gtk.FileFilterNew()
		if err != nil {
			log.Fatal("Failed to create filter:", err)
		}
		filter.AddPattern("*.db") // Filter for database files
		dialog.AddFilter(filter)

		response := dialog.Run()
		if response == gtk.RESPONSE_ACCEPT {
			// It is important to note that using := in a block of code where a variable has already
			// been declared at a higher visibility level will create a new local variable with the same name.
			// This new declaration will shadow the original variable within this code block.
			currentDatabasePath = dialog.GetFilename()
			getComboboxById("select-song-1", builder).RemoveAll()
			getComboboxById("select-song-2", builder).RemoveAll()
			getComboboxById("select-song-3", builder).RemoveAll()
			populateSongsAsync(builder)
			log.Println("Selected file:", currentDatabasePath)
		}
	})

	obj, err = builder.GetObject("exit-menu-item")
	if err != nil {
		log.Fatal("Could not find menu item 'Exit':", err)
	}

	exitMenuItem, ok := obj.(*gtk.MenuItem)
	if !ok {
		log.Fatal("Could not get menu item 'exit'")
	}

	exitMenuItem.Connect("activate", func() {
		gtk.MainQuit()
		log.Print("Exit")
	})

	win.Connect("destroy", func() {
		gtk.MainQuit()
		fmt.Println("Destroy")
	})

	return win
}

func loadSongsFromDatabase() ([]string, error) {

	// 	statement, _ := database.Prepare(`
	// 	CREATE TABLE songs (
	// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 		name TEXT NOT NULL,
	// 		type TEXT,
	// 		date DATE,
	// 		URL TEXT
	// 	);
	// `)
	// 	statement.Exec()

	// 	statement, _ = database.Prepare("INSERT INTO songs (name, type, date, URL) VALUES (?, ?, ?, ?);")
	// 	statement.Exec("Song 1", "Pop", "2023-03-01", "http://example.com/song1")
	// 	statement.Exec("Song 2", "Rock", "2023-03-02", "http://example.com/song2")
	// 	statement.Exec("Song 3", "Jazz", "2023-03-03", "http://example.com/song3")
	// 	statement.Exec("Song 4", "Classical", "2023-03-04", "http://example.com/song4")
	// 	statement.Exec("Song 5", "Electronic", "2023-03-05", "http://example.com/song5")

	database, err := sql.Open("sqlite", currentDatabasePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	defer database.Close()

	rows, err := database.Query("SELECT name FROM songs")
	if err != nil {
		return nil, fmt.Errorf("failed to query songs: %v", err)
	}
	defer rows.Close()

	var songs []string

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("failed to scan song: %v", err)
		}
		songs = append(songs, name) // Adding name of song to slice
	}

	// Check error after cycle
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through songs: %v", err)
	}

	return songs, nil
}

func populateSongsAsync(builder *gtk.Builder) {
	go func() {
		// Loading data from a database in a background thread
		songs, err := loadSongsFromDatabase()
		if err != nil {
			log.Println("Error loading songs", err)
			return
		}

		// Safe update of the GUI in the main thread
		glib.IdleAdd(func() bool {
			populateComboBox("select-song-1", builder, songs)
			populateComboBox("select-song-2", builder, songs)
			populateComboBox("select-song-3", builder, songs)
			return false // We return false so that the function is executed only once
		})
	}()
}

func populateComboBox(widgetID string, builder *gtk.Builder, songs []string) {

	comboBoxText := getComboboxById(widgetID, builder)

	for _, name := range songs {
		comboBoxText.AppendText(name)
	}
}

func getComboboxById(widgetID string, builder *gtk.Builder) *gtk.ComboBoxText {
	obj, err := builder.GetObject(widgetID)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not find '%s'", widgetID), err)
	}
	comboBoxText, ok := obj.(*gtk.ComboBoxText)

	if !ok {
		log.Fatal(fmt.Sprintf("Failed to get '%s' as ComboBoxText", widgetID))
	}

	return comboBoxText
}

func setColorWindow(win *gtk.Window) {
	// Creating a new CSS provider.
	cssProvider, err := gtk.CssProviderNew()
	if err != nil {
		log.Fatal("Failed to create CSS provider:", err)
	}

	cssPath, err := filepath.Abs("dark-theme.css")
	if err != nil {
		log.Fatal("Failed to get absolute path to CSS file:", err)
	}

	err = cssProvider.LoadFromPath(cssPath)
	if err != nil {
		log.Fatal("Failed to load CSS from file:", err)
	}

	// Adding CSS styles to the window.
	screen := win.GetScreen()
	gtk.AddProviderForScreen(screen, cssProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

}
