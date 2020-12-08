package window

import (
	"syscall"
	"unsafe"

	"github.com/JamesHovious/w32"
)

func MakeIntResource(id uint16) *uint16 {
	return (*uint16)(unsafe.Pointer(uintptr(id)))
}

func WndProc(hWnd w32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case w32.WM_DESTROY:
		w32.PostQuitMessage(0)
	default:
		return w32.DefWindowProc(hWnd, msg, wParam, lParam)
	}
	return 0
}

func Create() w32.HWND {
	hInstance := w32.GetModuleHandle("")
	var wcex w32.WNDCLASSEX
	wcex.Size = uint32(unsafe.Sizeof(wcex))
	wcex.Style = w32.CS_HREDRAW | w32.CS_VREDRAW
	wcex.WndProc = syscall.NewCallback(WndProc)
	wcex.ClsExtra = 0
	wcex.WndExtra = 0
	wcex.Instance = hInstance
	wcex.Icon = w32.LoadIcon(hInstance, MakeIntResource(w32.IDI_APPLICATION))
	wcex.Cursor = w32.LoadCursor(0, MakeIntResource(w32.IDC_ARROW))
	wcex.Background = w32.COLOR_WINDOW + 11
	lpszClassName := syscall.StringToUTF16Ptr("WNDclass")
	wcex.MenuName = nil

	wcex.ClassName = lpszClassName
	wcex.IconSm = w32.LoadIcon(hInstance, MakeIntResource(w32.IDI_APPLICATION))
	w32.RegisterClassEx(&wcex)

	hwnd := w32.CreateWindowEx(
		0, lpszClassName, syscall.StringToUTF16Ptr("DisVoice"),
		w32.WS_OVERLAPPEDWINDOW,
		w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, w32.CW_USEDEFAULT,
		0, 0, hInstance, nil)

	return hwnd
}

func Show(hWnd w32.HWND) {
	w32.ShowWindow(hWnd, w32.SW_SHOWDEFAULT)
	w32.UpdateWindow(hWnd)

	var msg w32.MSG
	for {
		if w32.GetMessage(&msg, 0, 0, 0) == 0 {
			break
		}
		w32.TranslateMessage(&msg)
		w32.DispatchMessage(&msg)
	}
}

type RECT struct {
	top    int
	bottom int
	left   int
	right  int
}

func SetSize(hWnd w32.HWND, height int, width int) {
	rcClient := w32.GetClientRect(hWnd)
	rcWind := w32.GetWindowRect(hWnd)
	x := (rcWind.Right - rcWind.Left) - rcClient.Right
	y := (rcWind.Bottom - rcWind.Top) - rcClient.Bottom
	w32.MoveWindow(hWnd, int(rcWind.Left), int(rcWind.Top), height+int(x), width+int(y), true)
}
