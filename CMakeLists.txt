cmake_minimum_required(VERSION 3.19.1)
project (disvoice)

# unicode
add_definitions(-DUNICODE)

# mingw compatibility
set(CMAKE_CXX_STANDARD_LIBRARIES "-static-libgcc -static-libstdc++ -lwsock32 -lws2_32 ${CMAKE_CXX_STANDARD_LIBRARIES}")
set(CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} -Wl,-Bstatic,--whole-archive -lwinpthread -Wl,--no-whole-archive")

# libraries
find_package(portaudio)

# compile
file(GLOB SOURCES "src/*.cpp")
add_executable(disvoice ${SOURCES})
target_link_libraries(disvoice portaudio)
