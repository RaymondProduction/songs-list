package gtk4r

/*
#cgo pkg-config: gtk4
#include <gtk/gtk.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// Ініціалізація GTK
func Init() {
	C.gtk_init(nil, nil)
}

// Builder структура, яка представляє GTK Builder
type Builder struct {
	builder *C.GtkBuilder
}

// BuilderNew створює новий екземпляр GTK Builder
func BuilderNew() (*Builder, error) {
	cBuilder := C.gtk_builder_new()
	if cBuilder == nil {
		return nil, fmt.Errorf("failed to create GTK Builder")
	}
	return &Builder{builder: cBuilder}, nil
}

// AddFromFile завантажує інтерфейс користувача з XML-файлу
func (b *Builder) AddFromFile(filename string) error {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	if C.gtk_builder_add_from_file(b.builder, cFilename, nil) == 0 {
		return fmt.Errorf("failed to load glade file")
	}
	return nil
}

// GetObject отримує об'єкт з білдера за ідентифікатором
func (b *Builder) GetObject(name string) (unsafe.Pointer, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	obj := C.gtk_builder_get_object(b.builder, cName)
	if obj == nil {
		return nil, fmt.Errorf("object with ID '%s' not found", name)
	}
	return unsafe.Pointer(obj), nil
}

// Window структура, яка представляє GTK Window
type Window struct {
	window *C.GtkWindow
}

// SetTitle встановлює заголовок вікна
func (w *Window) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.gtk_window_set_title(w.window, cTitle)
}

// ShowAll показує вікно та всі його віджети
func (w *Window) ShowAll() {
	C.gtk_widget_show_all((*C.GtkWidget)(unsafe.Pointer(w.window)))
}

// Main запускає головний цикл GTK
func Main() {
	C.gtk_main()
}
