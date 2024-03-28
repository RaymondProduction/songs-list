package main

import (
	"log"

	"./gtk4r"
)

func main() {
	gtk4r.Init(nil)

	builder, err := gtk4r.BuilderNew()
	if err != nil {
		log.Fatal("Error builder:", err)
	}

	err = builder.AddFromFile("main.glade")
	if err != nil {
		log.Fatal("Error when loading glade file:", err)
	}

	obj, err := builder.GetObject("main-window")
	if err != nil {
		log.Fatal("Error:", err)
	}

	win := (*gtk4r.Window)(obj)
	if win == nil {
		log.Fatal("Failed to get window")
	}

	win.SetTitle("Songs")

	win.ShowAll()

	gtk4r.Main()
}
