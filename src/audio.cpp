#include <portaudio.h>

#include "audio.h"
#include "error.h"
#include "common.h"

#define SAMPLE_RATE (44100)

typedef float SAMPLE;

typedef struct {
    int frameIndex;
    int maxFrameIndex;
    SAMPLE *recordedSample;
} paData;

void Audio::start()
{
    PaStreamParameters inputParameters,
        outputParameters;
    PaStream *stream;
    paData data;
    PaError err = paNoError;
    int i;
    int totalFrames;
    int numSamples;
    int numBytes;
    double average;
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
                               &data);         /*This is a pointer that will be passed to
                                                   your callback*/
}

/* This routine will be called by the PortAudio engine when audio is needed.
 * It may called at interrupt level on some machines so don't do anything
 * that could mess up the system like calling malloc() or free().
*/ 
static int patestCallback( const void *inputBuffer, void *outputBuffer,
                           unsigned long framesPerBuffer,
                           const PaStreamCallbackTimeInfo* timeInfo,
                           PaStreamCallbackFlags statusFlags,
                           void *userData )
{
    /* Cast data passed through stream to our structure. */
    void *data = (paData*)userData; 
    float *out = (float*)outputBuffer;
    unsigned int i;
    (void) inputBuffer; /* Prevent unused variable warning. */
    
    for( i=0; i<framesPerBuffer; i++ )
    {
         out++ = data->left_phase;  /* left */
         out++ = data->right_phase;  /* right */
        /* Generate simple sawtooth phaser that ranges between -1.0 and 1.0. */
        data->left_phase += 0.01f;
        /* When signal reaches top, drop back down. */
        if( data->left_phase >= 1.0f ) data->left_phase -= 2.0f;
        /* higher pitch so we can distinguish left and right. */
        data->right_phase += 0.03f;
        if( data->right_phase >= 1.0f ) data->right_phase -= 2.0f;
    }
    return 0;
}

void Audio::error(PaError err)
{
    Pa_Terminate();
    throw new Error(charToWString(Pa_GetErrorText(err)));
}
