#define UNICODE
#define _UNICODE

#ifdef __cplusplus
extern "C"
{
#endif

    void start();
    void stop();
    void *loadPlugin(const char *name);
    void startPlugin(void* plugin);
    void stopPlugin();
    void openPlugin(void* plugin);

#ifdef __cplusplus
}
#endif
