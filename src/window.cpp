#include "window.h"
#include "error.h"

#include <windows.h>

Window::Window(HINSTANCE instance, windowProc__ windowProc)
	: instance(instance), windowProc(windowProc) {}

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
		throw Error(L"Error creating window");
	}
}

void Window::show()
{
	ShowWindow(handle, SW_SHOW);

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
	windowClass.lpfnWndProc = windowProc;
	windowClass.hInstance = instance;
	windowClass.lpszClassName = className.c_str();

	if (!RegisterClassEx(&windowClass))
	{
		throw Error(L"Error registering window class");
	}
}

RECT Window::calcDimensions()
{
	RECT rect = {};

	rect.right = width;
	rect.bottom = height;

	if (!AdjustWindowRect(&rect, WS_SYSMENU | WS_CAPTION | WS_MINIMIZEBOX, false))
	{
		throw Error(L"Error calculating window size");
	}

	return rect;
}