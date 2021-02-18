#include "logger.hxx"
#include <iostream>

void Logger::info(std::string message, ...)
{
    std::cout << "INFO: " << message << std::endl;
}

void Logger::error(std::string message, ...)
{
    std::cout << "ERROR: " << message << std::endl;
}

void Logger::debug(std::string message, ...)
{
    std::cout << "DEBUG: " << message << std::endl;
}