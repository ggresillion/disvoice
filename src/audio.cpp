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
static int paCallback(const void *inputBuffer, void *outputBuffer,
                      unsigned long framesPerBuffer,
                      const PaStreamCallbackTimeInfo *timeInfo,
                      PaStreamCallbackFlags statusFlags,
                      void *userData)
{
    /* Cast data passed through stream to our structure. */
    SAMPLE *out = (SAMPLE *)outputBuffer;
    const SAMPLE *in = (const SAMPLE *)inputBuffer;

    for (int i = 0; i < framesPerBuffer; i++)
    {
        *out++ = in[i];
        *out++ = in[i];
    }

    if (plugin != NULL)
    {
        plugin->process(out, framesPerBuffer);
        return 0;
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
    }
    inputParameters.channelCount = 1; /* stereo input */
    inputParameters.sampleFormat = SAMPLE_TYPE;
    inputParameters.suggestedLatency = Pa_GetDeviceInfo(inputParameters.device)->defaultLowInputLatency;
    inputParameters.hostApiSpecificStreamInfo = NULL;

    outputParameters.device = Pa_GetDefaultOutputDevice(); /* default output device */
    if (outputParameters.device == paNoDevice)
    {
        fprintf(stderr, "Error: No default output device.\n");
    }
    outputParameters.channelCount = 2; /* stereo output */
    outputParameters.sampleFormat = SAMPLE_TYPE;
    outputParameters.suggestedLatency = Pa_GetDeviceInfo(outputParameters.device)->defaultLowOutputLatency;
    outputParameters.hostApiSpecificStreamInfo = NULL;

    /* Open an audio I/O stream. */
    err = Pa_OpenStream(&stream,
                        &inputParameters,
                        &outputParameters,
                        SAMPLE_RATE,
                        FRAMES_PER_BUFFER,
                        0,
                        paCallback,
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
