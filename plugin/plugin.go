package plugin

// #cgo CFLAGS: -I/home/guillaume/projects/disvoice/include
// #cgo pkg-config: portaudio-2.0
// #include "disvoice.hxx"
import "C"
import (
	"fmt"
	"unsafe"
)

type Plugin struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Filename string `json:"filename"`
	pointer  unsafe.Pointer
}

var plugin unsafe.Pointer

func Start() {
	fmt.Println("Starting ...")
	C.start()
}

func (p *Plugin) Load() {
	fmt.Println("Loading plugin " + p.Name)
	plugin = C.loadPlugin(C.CString(p.Filename))
	fmt.Println(plugin)
}

func (p *Plugin) Open() {
	C.openPlugin(p.pointer)
}

func (p *Plugin) StartPlugin() {
	C.startPlugin(p.pointer)
}

func StopPlugin() {
	C.stopPlugin()
}
