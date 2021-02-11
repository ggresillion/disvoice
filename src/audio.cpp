#include <portaudio.h>
#include <iostream>

#include "audio.h"
#include "error.h"
#include "config.h"
#include "logger.h"

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
static int paCallback(const void *inputBuffer, void *outputBuffer,
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
    const PaDeviceInfo *inputInfo;
    const PaDeviceInfo *outputInfo;
    int numChannels;

    err = Pa_Initialize();
    if (err != paNoError)
    {
        this->error(err);
    }
    Logger::info("Portaudio initialized");

    inputParameters.device = Pa_GetDefaultInputDevice(); /* default input device */
    inputInfo = Pa_GetDeviceInfo(inputParameters.device);

    outputParameters.device = Pa_GetDefaultOutputDevice(); /* default output device */
    outputInfo = Pa_GetDeviceInfo(outputParameters.device);

    numChannels = inputInfo->maxInputChannels < outputInfo->maxOutputChannels
                      ? inputInfo->maxInputChannels
                      : outputInfo->maxOutputChannels;

    inputParameters.channelCount = numChannels;
    inputParameters.sampleFormat = SAMPLE_TYPE;
    inputParameters.suggestedLatency = inputInfo->defaultHighInputLatency;
    inputParameters.hostApiSpecificStreamInfo = NULL;

    outputParameters.channelCount = numChannels;
    outputParameters.sampleFormat = SAMPLE_TYPE;
    outputParameters.suggestedLatency = outputInfo->defaultHighOutputLatency;
    outputParameters.hostApiSpecificStreamInfo = NULL;

    /* Open an audio I/O stream. */
    err = Pa_OpenStream(
        &stream,
        &inputParameters,
        &outputParameters,
        SAMPLE_RATE,
        FRAMES_PER_BUFFER,
        0,
        paCallback,
        NULL);
    if (err != paNoError)
        this->error(err);
    Logger::info("Portaudio stream created");

    err = Pa_StartStream(stream);
    if (err != paNoError)
        this->error(err);
    Logger::info("Portaudio stream started");
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
    throw new Error(Pa_GetErrorText(err));
}
