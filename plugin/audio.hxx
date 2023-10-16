#pragma once

#include <portaudio.h>

#include "vst_plugin.hxx"

class Audio
{
public:
    static void setPlugin(VstPlugin *plugin);
    void start();
    void stop();

private:
    PaStream *stream;
    void error(PaError err);
};