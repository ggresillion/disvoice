#include "iostream"
#include "vst_plugin.hxx"

class Effect
{
public:
    Effect();
    Effect(std::string name, VstPlugin *plugin);

    std::string getName();
    VstPlugin *getPlugin();

private:
    std::string name;
    VstPlugin *plugin;
};