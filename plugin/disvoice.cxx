#include <iostream>
#include <sstream>

#include "disvoice.hxx"
#include "window.hxx"
#include "logger.hxx"
#include "config.hxx"
#include "audio.hxx"

const std::wstring ConfigPath = L"MiniVST.cfg";
const std::wstring WindowClass = L"MiniVST";
const std::wstring WindowTitle = L"MiniVST v0.0.1 �Severin v. W.";
const std::string PluginPath = "plugins/";

Audio *audio;
Window *window;

void start()
{
    audio = new Audio();
    audio->start();
}

void stop()
{
    audio->stop();
}

void *loadPlugin(const char *name)
{
    VstPlugin *plugin;
    std::string str(name);
    plugin = new VstPlugin(PluginPath + name);
    plugin->start(SAMPLE_RATE);
    Logger::info("Started plugin");
    int width, height;
    plugin->getEditorRect(width, height);

    window = new Window(GetModuleHandle(NULL));
    Logger::info("Got plugin rect");
    window->setClassName(WindowClass);
    window->setTitle(WindowTitle);
    window->setSize(width, height);
    window->create();

    plugin->openEditor(window->getHandle());

    window->show();

    return static_cast<void *>(plugin);
}

void startPlugin(void *plugin)
{
    audio->setPlugin(static_cast<VstPlugin *>(plugin));
}

void stopPlugin()
{
    audio->setPlugin(nullptr);
}

void openPlugin(void *p)
{
    VstPlugin *plugin = static_cast<VstPlugin *>(p);
    Logger::info("Opening plugin");

    window->show();
    Logger::info("Failed to open");
}
