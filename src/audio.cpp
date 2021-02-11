#include <portaudio.h>
#include <iostream>

#include "audio.h"
#include "error.h"
#include "common.h"
#include "config.h"

using namespace std;

typedef float SAMPLE;

typedef struct
{
    int frameIndex;
    int maxFrameIndex;
    SAMPLE *recordedSample;
} paData;

VstPlugin *plugin;

/* This routine will be called by the PortAudio engine when audio is needed.
 * It may called at interrupt level on some machines so don't do anything
 * that could mess up the system like calling malloc() or free().
*/
static int patestCallback(const void *inputBuffer, void *outputBuffer,
                          unsigned long framesPerBuffer,
                          const PaStreamCallbackTimeInfo *timeInfo,
                          PaStreamCallbackFlags statusFlags,
                          void *userData)
{
    /* Cast data passed through stream to our structure. */
    float *in = (float *)inputBuffer;
    float *out = (float *)outputBuffer;

    if (plugin != NULL)
    {
        plugin->process(in, out, framesPerBuffer);
        return 0;
    }

    for (unsigned int i = 0; i < framesPerBuffer; i++)
    {
        *out++ = *in++; /* left */
        *out++ = *in++; /* right */
    }
    return 0;
}

void Audio::start()
{
    PaStreamParameters inputParameters, outputParameters;
    PaStream *stream;
    PaError err;

    err = Pa_Initialize();
    if (err != paNoError)
        this->error(err);

    /* Open an audio I/O stream. */
    err = Pa_OpenDefaultStream(&stream,
                               0,         /* no input channels */
                               2,         /* stereo output */
                               paFloat32, /* 32 bit floating point output */
                               SAMPLE_RATE,
                               256,            /* frames per buffer, i.e. the number
                                                   of sample frames that PortAudio will
                                                   request from the callback. Many apps
                                                   may want to use
                                                   paFramesPerBufferUnspecified, which
                                                   tells PortAudio to pick the best,
                                                   possibly changing, buffer size.*/
                               patestCallback, /* this is your callback function */
                               NULL);         /*This is a pointer that will be passed to
                                                   your callback*/
    if (err != paNoError)
        this->error(err);
    cout << "Portaudio stream created" << endl;
    err = Pa_StartStream(stream);
    if (err != paNoError)
        this->error(err);
    cout << "Portaudio stream started" << endl;
}

void Audio::stop()
{
    PaError err = paNoError;
    err = Pa_CloseStream(stream);
    if (err != paNoError)
        this->error(err);
    err = Pa_Terminate();
    if (err != paNoError)
        this->error(err);
}

void Audio::setPlugin(VstPlugin *p)
{
    plugin = p;
}

void Audio::error(PaError err)
{
    Pa_Terminate();
    throw new Error(charToWString(Pa_GetErrorText(err)));
}
