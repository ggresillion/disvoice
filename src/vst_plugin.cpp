/*
 * Copyright (C) 2018 Medusalix
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
#include <windows.h>
#include <aeffect.h>
#include <iostream>

#include "vst_plugin.h"
#include "error.h"

#define SAMPLE_RATE (44100)

using VstMain = AEffect *(*)(audioMasterCallback callback);

VstPlugin::VstPlugin(const std::string &path) : path(path)
{
	HMODULE library = LoadLibraryA(path.c_str());

	if (library == nullptr)
	{
		std::cout << GetLastError() << std::endl;
		throw Error(L"Error loading plugin");
	}

	FARPROC mainAddress = GetProcAddress(library, "VSTPluginMain");

	if (mainAddress == nullptr)
	{
		mainAddress = GetProcAddress(library, "main");
	}

	if (mainAddress == nullptr)
	{
		throw Error(L"Plugin entry point not found");
	}

	effect = reinterpret_cast<VstMain>(mainAddress)(hostCallback);

	if (effect == nullptr)
	{
		throw Error(L"Plugin initialization failed");
	}

	if (effect->magic != kEffectMagic)
	{
		throw Error(L"Plugin magic number mismatch");
	}

	if (effect->numOutputs > 2)
	{
		throw Error(L"Plugin has more than 2 output channels");
	}

	size_t fileIndex = path.find_last_of(L'/');

	if (fileIndex == std::wstring::npos)
	{
		fileIndex = path.find_last_of(L'\\');
	}

	size_t extIndex = path.find_last_of(L'.');

	name = path.substr(fileIndex + 1, extIndex - fileIndex - 1);

}

std::string VstPlugin::getName() const
{
	return name;
}

void VstPlugin::start(int sampleRate)
{
	// Block size is sample rate [kHz] * channels
	int blockSize = sampleRate / 100 * 2;

	input.emplace_back(blockSize);
	input.emplace_back(blockSize);
	output.emplace_back(blockSize);
	output.emplace_back(blockSize);

	rawInput[0] = input[0].data();
	rawInput[1] = input[1].data();
	rawOutput[0] = output[0].data();
	rawOutput[1] = output[1].data();

	effect->dispatcher(effect, effOpen, 0, 0, nullptr, 0);
	effect->dispatcher(effect, effSetSampleRate, 0, 0, nullptr, static_cast<float>(sampleRate));
	effect->dispatcher(effect, effSetBlockSize, 0, blockSize, nullptr, 0);
	effect->dispatcher(effect, effMainsChanged, 0, 1, nullptr, 0);

}

void VstPlugin::getEditorRect(int &width, int &height)
{
	ERect **rect = new ERect *;
	effect->dispatcher(effect, effEditGetRect, 0, 0, rect, 0);

	width = (*rect)->right - (*rect)->left;
	height = (*rect)->bottom - (*rect)->top;

	delete rect;
}

void VstPlugin::openEditor(void *windowHandle)
{
	effect->dispatcher(effect, effEditOpen, 0, 0, windowHandle, 0);
}

void VstPlugin::process(float *inBuffer, float *outBuffer, int size)
{
	int framesPerBuffer = size / 2;
	// Move both channels into separate arrays
	for (int i = 0; i < framesPerBuffer; i++)
	{
		input[0][i] = inBuffer[i * 2];
		input[1][i] = inBuffer[i * 2 + 1];
	}

	if (effect->flags & effFlagsCanReplacing)
	{
		effect->processReplacing(effect, rawInput, rawOutput, framesPerBuffer);
	}

	else
	{
		effect->__processDeprecated(effect, rawInput, rawOutput, framesPerBuffer);
	}

	// Move output back into one array
	for (int i = 0; i < framesPerBuffer; i++)
	{
		outBuffer[i * 2] = output[0][i];
		outBuffer[i * 2 + 1] = output[1][i];
	}
}

VstIntPtr VstPlugin::hostCallback(
	AEffect *effect,
	VstInt32 opcode,
	VstInt32 index,
	VstIntPtr value,
	void *ptr,
	float opt)
{
	switch (opcode)
	{
	case audioMasterVersion:
		// VST 2.4
		return 2400;

	default:
		return 0;
	}
}
