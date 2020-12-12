cmake_minimum_required(VERSION 3.19.1)
project (disvoice)

add_definitions(-DUNICODE)

# compile a win32 application
add_compile_options(-Wl,-mwindows)

# link statically some gcc dlls
set(CMAKE_CXX_STANDARD_LIBRARIES "-static-libgcc -static-libstdc++ -lwsock32 -lws2_32 ${CMAKE_CXX_STANDARD_LIBRARIES}")
set(CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} -Wl,-Bstatic,--whole-archive -lwinpthread -Wl,--no-whole-archive")

# the name of the target operating system
set(CMAKE_SYSTEM_NAME Windows)

# which compilers to use for C and C++
set(CMAKE_C_COMPILER x86_64-w64-mingw32-gcc)
set(CMAKE_CXX_COMPILER x86_64-w64-mingw32-g++)

# includes
include_directories(include)
include_directories(/usr/x86_64-w64-mingw32/include)

# compile
file(GLOB SOURCES "src/*.cpp")
add_executable(disvoice ${SOURCES})

# link against libraries
target_link_libraries(disvoice /usr/local/lib/libportaudio-2.dll)
