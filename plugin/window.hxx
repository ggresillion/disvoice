#include <string>
#ifndef UNICODE
#define UNICODE
#define UNICODE_WAS_UNDEFINED
#endif

#include <windows.h>

#ifdef UNICODE_WAS_UNDEFINED
#undef UNICODE
#endif

struct HINSTANCE__;
struct HWND__;

using HINSTANCE = HINSTANCE__*;
using HWND = HWND__*;

using UINT = unsigned int;
using LONG = long;
using RECT = struct tagRECT;

typedef LRESULT CALLBACK (*windowProc__) (HWND hwnd, UINT uMsg, WPARAM wParam, LPARAM lParam);

class Window
{
public:
	Window(HINSTANCE instance);

	void create();
	void show();

	void setClassName(const std::wstring &className);
	void setTitle(const std::wstring &title);
	void setIcon(int icon);
	void setSize(int width, int height);

	HWND getHandle() const;

private:
	HINSTANCE instance;
	windowProc__ windowProc;
	HWND handle = nullptr;

	std::wstring className;
	std::wstring title;

	int width = 0, height = 0;

	void initWindowClass();
	RECT calcDimensions();
};