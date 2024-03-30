package gtk4r

/*
#cgo pkg-config: gtk4
#include <gtk/gtk.h>
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

type Application struct {
	app *C.GtkApplication
}

func NewApplication(appID string, flags uint) *Application {
	cAppID := C.CString(appID)
	defer C.free(unsafe.Pointer(cAppID))

	return &Application{
		app: C.gtk_application_new(cAppID, C.GApplicationFlags(flags)),
	}
}

func (a *Application) Run() int {
	defer C.g_object_unref(C.gpointer(a.app))
	return int(C.g_application_run((*C.GApplication)(unsafe.Pointer(a.app)), 0, nil))
	//return int(C.g_application_run(C.GApplication(a.app), 0, nil))
	//return int(C.g_application_run(C.G_APPLICATION(a.app), 0, nil))
	//return int(C.g_application_run((*C.GApplication)(a.app), 0, nil))
}

type ActivateFunc func(app *Application)

var activateFunc ActivateFunc

//export goActivateFunc
func goActivateFunc(app *C.GtkApplication, userData C.gpointer) {
	if activateFunc != nil {
		activateFunc(&Application{app: app})
	}
}

func (a *Application) ConnectActivate(fn ActivateFunc) {
	activateFunc = fn
	C.g_signal_connect_data(C.gpointer(a.app), C.CString("activate"),
		(C.GCallback)(C.goActivateFunc), nil, nil, 0)
}

type Window struct {
	window *C.GtkWindow
}

func (a *Application) NewWindow() *Window {
	cWindow := C.gtk_application_window_new(a.app)
	window := (*C.GtkWindow)(unsafe.Pointer(cWindow))
	return &Window{window: window}
}

func (w *Window) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.gtk_window_set_title(w.window, cTitle)
}

func (w *Window) SetDefaultSize(width, height int) {
	C.gtk_window_set_default_size(w.window, C.int(width), C.int(height))
}
