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

    inputParameters.device = Pa_GetDefaultInputDevice(); /* default input device */
    if (inputParameters.device == paNoDevice)
    {
        fprintf(stderr, "Error: No default input device.\n");
        this->error(err);
    }
    inputParameters.channelCount = 2; /* stereo input */
    inputParameters.sampleFormat = PA_SAMPLE_TYPE;
    inputParameters.suggestedLatency = Pa_GetDeviceInfo(inputParameters.device)->defaultLowInputLatency;
    inputParameters.hostApiSpecificStreamInfo = NULL;

    outputParameters.device = Pa_GetDefaultOutputDevice(); /* default output device */
    if (outputParameters.device == paNoDevice)
    {
        fprintf(stderr, "Error: No default output device.\n");
        this->error(err);
    }
    outputParameters.channelCount = 2; /* stereo output */
    outputParameters.sampleFormat = PA_SAMPLE_TYPE;
    outputParameters.suggestedLatency = Pa_GetDeviceInfo(outputParameters.device)->defaultLowOutputLatency;
    outputParameters.hostApiSpecificStreamInfo = NULL;
    cout << "Portaudio initialized" << endl;
    /* Open an audio I/O stream. */
    err = Pa_OpenStream(
        &stream,
        &inputParameters,
        &outputParameters,
        SAMPLE_RATE,
        FRAMES_PER_BUFFER,
        0, /* paClipOff, */ /* we won't output out of range samples so don't bother clipping them */
        patestCallback,
        NULL);
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
