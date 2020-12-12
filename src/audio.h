#pragma once

#include <portaudio.h>

class Audio {
    public:
        void start();
    private:
        void error(PaError err);
};