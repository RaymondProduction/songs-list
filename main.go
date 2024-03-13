package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/glebarez/go-sqlite"
	"github.com/gotk3/gotk3/gtk"
)

func main() {

	gtk.Init(nil)
	win := initGTKWindow()

	go func() {
		test()
	}()

	win.ShowAll()

	gtk.Main()

}

func initGTKWindow() *gtk.Window {

	// Create builder
	builder, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Error bulder:", err)
	}

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

		response := dialog.Run()
		if response == gtk.RESPONSE_ACCEPT {
			filename := dialog.GetFilename()
			log.Println("Selected file:", filename)
		}
	})

	win.Connect("destroy", func() {
		//gtk.Quit()
		fmt.Println("Destroy")
	})

	return win
}

func test() {
	database, _ := sql.Open("sqlite", "./songs.db")
	statement, _ := database.Prepare(`
	CREATE TABLE songs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		type TEXT,
		date DATE,
		URL TEXT
	);
`)
	statement.Exec()

	statement, _ = database.Prepare("INSERT INTO songs (name, type, date, URL) VALUES (?, ?, ?, ?);")
	statement.Exec("Song 1", "Pop", "2023-03-01", "http://example.com/song1")
	statement.Exec("Song 2", "Rock", "2023-03-02", "http://example.com/song2")
	statement.Exec("Song 3", "Jazz", "2023-03-03", "http://example.com/song3")
	statement.Exec("Song 4", "Classical", "2023-03-04", "http://example.com/song4")
	statement.Exec("Song 5", "Electronic", "2023-03-05", "http://example.com/song5")

	rows, _ := database.Query("SELECT id, name, date FROM songs")

	var id int
	var name, date string

	for rows.Next() {
		err := rows.Scan(&id, &name, &date)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %s %s\n", id, name, date)
	}
}
