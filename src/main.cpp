#include <windows.h>
#include <string>
#include <iostream>
#include <functional>
#include <string>
using namespace std;

#include "error.h"
#include "vst_plugin.h"
#include "audio.h"
#include "window.h"

#ifndef UNICODE
#define UNICODE
#endif

using namespace std::placeholders;

const std::wstring ConfigPath = L"MiniVST.cfg";
const std::wstring WindowClass = L"MiniVST";
const std::wstring WindowTitle = L"MiniVST v0.0.1 �Severin v. W.";

LRESULT CALLBACK WindowProc(HWND hwnd, UINT uMsg, WPARAM wParam, LPARAM lParam);

int main()
{
   return WinMain(GetModuleHandle(NULL), NULL, GetCommandLineA(), SW_SHOWNORMAL);
}

int WINAPI WinMain(HINSTANCE hInstance, HINSTANCE, LPSTR pCmdLine, int nCmdShow)
{
    try
    {
        VstPlugin *plugin = new VstPlugin("../plugins/TAL-Reverb-4-64.dll");
        plugin->start(44100);

        Audio *audio = new Audio();
        audio->start();
        audio->setPlugin(plugin);

        int width, height;
        plugin->getEditorRect(width, height);

        Window *window = new Window(hInstance, WindowProc);
        window->setClassName(WindowClass);
        window->setTitle(WindowTitle);
        window->setSize(width, height);
        window->create();

        plugin->openEditor(window->getHandle());

        window->show();

        audio->stop();
    }
    catch (Error err)
    {
        cout << err.getMessage() << "\n";
    }

    return 0;
}

LRESULT CALLBACK WindowProc(HWND hwnd, UINT uMsg, WPARAM wParam, LPARAM lParam)
{
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
    }
    return DefWindowProc(hwnd, uMsg, wParam, lParam);
}