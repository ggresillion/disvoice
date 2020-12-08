package plugin

/*
typedef char VstInt8;				///< 8 bit integer type

#ifdef WIN32
	typedef short VstInt16;			///< 16 bit integer type
	typedef int VstInt32;			///< 32 bit integer type
	typedef __int64 VstInt64;		///< 64 bit integer type
#else
	#include <stdint.h>
	typedef int16_t VstInt16;		///< 16 bit integer type
	typedef int32_t VstInt32;		///< 32 bit integer type
	typedef int64_t VstInt64;		///< 64 bit integer type
#endif

typedef struct {
	VstInt16 top;		///< top coordinate
	VstInt16 left;		///< left coordinate
	VstInt16 bottom;	///< bottom coordinate
	VstInt16 right;		///< right coordinate
} ERect;
*/
import "C"
import (
	"fmt"
	"ggresillion/disvoice/config"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"unsafe"

	"github.com/JamesHovious/w32"
	"pipelined.dev/audio/vst2"
	"pipelined.dev/signal"
)

type Plugin struct {
	vst vst2.Plugin
}

var bin vst2.FloatBuffer
var bout vst2.FloatBuffer
var buffer signal.Floating

func InitPlugin(path string) Plugin {
	// Open VST library. Library contains a reference to
	// OS-specific handle, that needs to be freed with Close.
	vst, err := vst2.Open(pluginPath(path))
	if err != nil {
		log.Panicf("failed to open VST library: %v", err)
	}

	// Load VST plugin with example callback.
	plugin := vst.Load(PrinterHostCallback("Received opcode"))

	return Plugin{*plugin}
}

func (p Plugin) ProcessAudio(in, out []float32) {

	signal.WriteFloat32(in, buffer)

	bin.CopyFrom(buffer)

	p.vst.ProcessFloat(bin, bout)

	bout.CopyTo(buffer)

	signal.ReadFloat32(buffer, out)

}

// PrinterHostCallback returns closure that prints received opcode with
// provided prefix. This technique allows to provide callback with any
// context needed.
func PrinterHostCallback(prefix string) vst2.HostCallbackFunc {
	return func(code vst2.HostOpcode, _ vst2.Index, _ vst2.Value, _ vst2.Ptr, _ vst2.Opt) vst2.Return {
		fmt.Printf("%s: %v\n", prefix, code)
		switch code {
		case vst2.HostVersion:
			// VST 2.4
			return 2400

		case vst2.HostGetCurrentProcessLevel:
			return 1

		default:
			return 0
		}
	}
}

func initBuffers(config config.AudioConfig) {
	bin = vst2.NewFloatBuffer(config.Channels, config.BufferLength)
	bout = vst2.NewFloatBuffer(config.Channels, config.BufferLength)
	buffer = signal.Allocator{
		Channels: config.Channels,
		Length:   config.BufferLength,
		Capacity: config.BufferLength,
	}.Float32()
}

func (p Plugin) LoadSettings() {

	data, err := ioutil.ReadFile("save.bin")
	chk(err)
	fmt.Println(data[:])
	s := C.CString(string(data[:]))
	p.vst.Dispatch(vst2.EffSetChunk, 0, 4, vst2.Ptr(s), 0)
}

func (p Plugin) SaveSettings() {

	var chunk *C.char
	chunkSize := p.vst.Dispatch(vst2.EffGetChunk, 0, 0, vst2.Ptr(&chunk), 0)
	bytes := C.GoBytes(unsafe.Pointer(&chunk), C.int(unsafe.Sizeof(C.int(chunkSize))))
	fmt.Println(bytes)
	ioutil.WriteFile("save.bin", bytes, os.ModeExclusive)
}

func (p Plugin) Configure(config config.AudioConfig) {

	// Set sample rate in Hertz.
	p.vst.SetSampleRate(config.SampleRate)
	// Set channels information.
	p.vst.SetSpeakerArrangement(
		&vst2.SpeakerArrangement{
			Type:        vst2.SpeakerArrMono,
			NumChannels: int32(config.Channels),
		},
		&vst2.SpeakerArrangement{
			Type:        vst2.SpeakerArrMono,
			NumChannels: int32(config.Channels),
		},
	)
	// Set buffer size.
	p.vst.SetBufferSize(config.BufferLength)

	initBuffers(config)

}

func (p Plugin) Start() {
	// Start the plugin.
	p.vst.Start()
}

// pluginPath returns a path to OS-specific plugin. It will panic if OS is
// not supported.
func pluginPath(path string) string {
	absPath, _ := filepath.Abs(path)
	return absPath
}

type rectangle struct {
	top    int
	bottom int
	left   int
	right  int
}

/*
	GUI
*/

func (p Plugin) GetEditorRect() (int, int) {

	var rect *C.ERect

	p.vst.Dispatch(vst2.EffEditGetRect, 0, 0, vst2.Ptr(&rect), 0)

	return int(rect.right - rect.left), int(rect.bottom - rect.top)
}

func (p Plugin) OpenGui(handle w32.HWND) {
	p.vst.Dispatch(vst2.EffEditOpen, 0, 0, vst2.Ptr(handle), 0)
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
