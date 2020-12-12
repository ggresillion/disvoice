#include <locale>
#include <codecvt>
#include <string>
#include <stdio.h>

using namespace std;

wstring charToWString(const char* data) {
       static std::wstring_convert<std::codecvt_utf8_utf16<wchar_t>> converter;
    return wstring(converter.from_bytes(data));
}