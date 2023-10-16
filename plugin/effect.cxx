#include "effect.hxx"

Effect::Effect(std::string name, VstPlugin *plugin) : name(name), plugin(plugin) {}

std::string Effect::getName()
{
    return name;
}

VstPlugin *Effect::getPlugin()
{
    return plugin;
}