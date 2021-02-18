#include <string>

class Logger
{
public:
    static void info(std::string message, ...);
    static void error(std::string message, ...);
    static void debug(std::string message, ...);
};