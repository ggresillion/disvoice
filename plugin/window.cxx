#include "window.hxx"
#include "error.hxx"

#include <windows.h>

LRESULT CALLBACK WindowProc(HWND hwnd, UINT uMsg, WPARAM wParam, LPARAM lParam)
{
	std::cout << uMsg << std::endl;
	switch (uMsg)
	{
	case WM_DESTROY:
		PostQuitMessage(0);
		return 0;

	case WM_PAINT:
	{
		PAINTSTRUCT ps;
		HDC hdc = BeginPaint(hwnd, &ps);

		FillRect(hdc, &ps.rcPaint, (HBRUSH)(COLOR_WINDOW + 1));

		EndPaint(hwnd, &ps);
	}
		return 0;
	case WM_SETFOCUS:
	{
		return 0;
	}
	}
	return DefWindowProc(hwnd, uMsg, wParam, lParam);
}

Window::Window(HINSTANCE instance)
{
	instance = GetModuleHandle(NULL);
}

void Window::create()
{
	initWindowClass();
	RECT rect = calcDimensions();

	handle = CreateWindow(
		className.c_str(),
		title.c_str(),
		WS_SYSMENU | WS_CAPTION | WS_MINIMIZEBOX,
		CW_USEDEFAULT,
		CW_USEDEFAULT,
		rect.right - rect.left,
		rect.bottom - rect.top,
		nullptr,
		nullptr,
		instance,
		nullptr);

	if (handle == nullptr)
	{
		throw Error("Error creating window");
	}
}

void Window::show()
{
	std::cout << handle << std::endl;
	ShowWindow(handle, SW_SHOW);

	std::cout << GetLastError() << std::endl;

	MSG message;

	while (GetMessage(&message, nullptr, 0, 0) != 0)
	{
		TranslateMessage(&message);
		DispatchMessage(&message);
	}
}

void Window::setClassName(const std::wstring &className)
{
	this->className = className;
}

void Window::setTitle(const std::wstring &title)
{
	if (handle == nullptr)
	{
		this->title = title;
	}

	else
	{
		SetWindowText(handle, title.c_str());
	}
}

void Window::setSize(int width, int height)
{
	this->width = width;
	this->height = height;
}

HWND Window::getHandle() const
{
	return handle;
}

void Window::initWindowClass()
{

	WNDCLASSEX windowClass = {};

	windowClass.cbSize = sizeof(WNDCLASSEX);
	windowClass.lpfnWndProc = WindowProc;
	windowClass.hInstance = instance;
	windowClass.lpszClassName = className.c_str();

	if (!RegisterClassEx(&windowClass))
	{
		throw Error("Error registering window class");
	}
}

RECT Window::calcDimensions()
{
	RECT rect = {};

	rect.right = width;
	rect.bottom = height;

	if (!AdjustWindowRect(&rect, WS_SYSMENU | WS_CAPTION | WS_MINIMIZEBOX, false))
	{
		throw Error("Error calculating window size");
	}

	return rect;
}