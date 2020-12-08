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

// var bin vst2.DoubleBuffer
// var bout vst2.DoubleBuffer
// var buffer signal.Floating

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

func (p Plugin) ProcessAudio(_, bout []float32) {

	// signal.WriteFloat32(in, buffer)

	// bin.CopyFrom(buffer)

	// p.vst.ProcessDouble(bin, bout)

	// bout.CopyTo(buffer)

	// signal.ReadFloat32(buffer, out)

	data := [][]float64{
		{
			-0.0027160645,
			-0.0039978027,
			-0.0071411133,
			-0.0065307617,
			0.0038757324,
			0.021972656,
			0.041229248,
			0.055511475,
			0.064971924,
			0.07342529,
			0.08300781,
			0.092681885,
			0.10070801,
			0.110809326,
			0.12677002,
			0.15231323,
			0.19058228,
			0.24459839,
			0.3140869,
			0.38861084,
			0.44683838,
			0.47177124,
			0.46643066,
			0.45007324,
			0.4449768,
			0.45724487,
			0.47451782,
			0.48321533,
			0.47824097,
			0.46679688,
			0.45999146,
			0.46765137,
		},
		{
			-0.0027160645,
			-0.0039978027,
			-0.0071411133,
			-0.0065307617,
			0.0038757324,
			0.021972656,
			0.041229248,
			0.055511475,
			0.064971924,
			0.07342529,
			0.08300781,
			0.092681885,
			0.10070801,
			0.110809326,
			0.12677002,
			0.15231323,
			0.19058228,
			0.24459839,
			0.3140869,
			0.38861084,
			0.44683838,
			0.47177124,
			0.46643066,
			0.45007324,
			0.4449768,
			0.45724487,
			0.47451782,
			0.48321533,
			0.47824097,
			0.46679688,
			0.45999146,
			0.46765137,
		},
	}

	buffer := signal.Allocator{
		Channels: len(data),
		Length:   len(data[0]),
		Capacity: len(data[0]),
	}.Float64()
	signal.WriteStripedFloat64(data, buffer)

	// To process data with plugin, we need to use VST2 buffers.
	// It's needed because VST SDK was written in C and expected
	// memory layout differs from Golang slices.
	// We need two buffers for input and output.
	in := vst2.NewDoubleBuffer(buffer.Channels(), buffer.Length())
	defer in.Free()
	out := vst2.NewDoubleBuffer(buffer.Channels(), buffer.Length())
	defer out.Free()

	// Fill input with data values.
	in.CopyFrom(buffer)

	// Process data.
	p.vst.ProcessDouble(in, out)
	// Copy processed data.
	out.CopyTo(buffer)

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
	// bin = vst2.NewDoubleBuffer(config.Channels, config.BufferLength)
	// bout = vst2.NewDoubleBuffer(config.Channels, config.BufferLength)
	// buffer = signal.Allocator{
	// 	Channels: config.Channels,
	// 	Length:   config.BufferLength,
	// 	Capacity: config.BufferLength,
	// }.Float32()
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
	p.vst.SetSampleRate(44100)
	// Set channels information.
	p.vst.SetSpeakerArrangement(
		&vst2.SpeakerArrangement{
			Type:        vst2.SpeakerArrMono,
			NumChannels: int32(1),
		},
		&vst2.SpeakerArrangement{
			Type:        vst2.SpeakerArrMono,
			NumChannels: int32(1),
		},
	)
	// Set buffer size.
	p.vst.SetBufferSize(32)

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
