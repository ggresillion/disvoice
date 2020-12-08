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
	"log"
	"path/filepath"

	"github.com/JamesHovious/w32"
	"pipelined.dev/audio/vst2"
	"pipelined.dev/signal"
)

var bin vst2.FloatBuffer
var bout vst2.FloatBuffer
var buffer signal.Floating
var plugin *vst2.Plugin

func InitPlugin(path string) {
	// Open VST library. Library contains a reference to
	// OS-specific handle, that needs to be freed with Close.
	vst, err := vst2.Open(pluginPath(path))
	if err != nil {
		log.Panicf("failed to open VST library: %v", err)
	}

	// Load VST plugin with example callback.
	plugin = vst.Load(PrinterHostCallback("Received opcode"))
}

func ProcessAudio(in, out []float32) {

	signal.WriteFloat32(in, buffer)

	bin.CopyFrom(buffer)

	plugin.ProcessFloat(bin, bout)

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

// func (p Plugin) LoadSettings() {

// 	data, err := ioutil.ReadFile("save.bin")
// 	chk(err)
// 	fmt.Println(data[:])
// 	s := C.CString(string(data[:]))
// 	p.vst.Dispatch(vst2.EffSetChunk, 0, 4, vst2.Ptr(s), 0)
// }

// func (p Plugin) SaveSettings() {

// 	var chunk *C.char
// 	chunkSize := p.vst.Dispatch(vst2.EffGetChunk, 0, 0, vst2.Ptr(&chunk), 0)
// 	bytes := C.GoBytes(unsafe.Pointer(&chunk), C.int(unsafe.Sizeof(C.int(chunkSize))))
// 	fmt.Println(bytes)
// 	ioutil.WriteFile("save.bin", bytes, os.ModeExclusive)
// }

func Configure(config config.AudioConfig) {

	// Set sample rate in Hertz.
	plugin.SetSampleRate(config.SampleRate)
	// Set channels information.
	plugin.SetSpeakerArrangement(
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
	plugin.SetBufferSize(config.BufferLength)

	bin = vst2.NewFloatBuffer(config.Channels, config.BufferLength)
	bout = vst2.NewFloatBuffer(config.Channels, config.BufferLength)
	buffer = signal.Allocator{
		Channels: config.Channels,
		Length:   config.BufferLength,
		Capacity: config.BufferLength,
	}.Float32()

}

func Start() {
	// Start the plugin.
	plugin.Start()
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

func GetEditorRect() (int, int) {

	var rect *C.ERect

	plugin.Dispatch(vst2.EffEditGetRect, 0, 0, vst2.Ptr(&rect), 0)

	width := int(rect.right - rect.left)
	height := int(rect.bottom - rect.top)

	return width, height
}

func OpenGui(handle w32.HWND) {
	plugin.Dispatch(vst2.EffEditOpen, 0, 0, vst2.Ptr(handle), 0)
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
