package main

import (
	"fmt"
	"gtk4r/gtk4r"
	"unsafe"
)

//export printHello
func printHello(widget unsafe.Pointer, data unsafe.Pointer) {
	fmt.Println("Hello World")
}

func main() {
	app := gtk4r.NewApplication("org.gtk.example", 0)

	app.ConnectActivate(func(app *gtk4r.Application) {
		fmt.Println("Application Activated")
	})

	status := app.Run()
	fmt.Println("Application exited with status:", status)

}
