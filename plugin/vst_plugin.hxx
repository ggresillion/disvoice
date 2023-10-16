#pragma once

#include <string>
#include <vector>
#include "../include/aeffect.h"

struct AEffect;

class VstPlugin
{
public:
	VstPlugin(const std::string &path);

	std::string getName() const;

	void loadSettings();
	void saveSettings();

	void start(int sampleRate);

	void getEditorRect(int &width, int &height);
	void openEditor(void *windowHandle);

	void process(float *in, float* out, int framesPerBuffer);

private:
	AEffect *effect;
	std::string name;
	std::string path;

	std::vector<std::vector<float>> input, output;
	float *rawInput[2], *rawOutput[2];

	static VstIntPtr hostCallback(
		AEffect *effect,
		VstInt32 opcode,
		VstInt32 index,
		VstIntPtr value,
		void *ptr,
		float opt);
};
