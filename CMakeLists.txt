cmake_minimum_required(VERSION 3.19.1)
project (disvoice)

add_definitions(-DUNICODE)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Find includes in corresponding build directories
set(CMAKE_INCLUDE_CURRENT_DIR ON)
# Instruct CMake to run moc automatically when needed.
set(CMAKE_AUTOMOC ON)


# compile a win32 application
add_compile_options(-Wl,-mwindows)


# the name of the target operating system
set(CMAKE_SYSTEM_NAME Windows)

# which compilers to use for C and C++
set(CMAKE_C_COMPILER x86_64-w64-mingw32-gcc)
set(CMAKE_CXX_COMPILER x86_64-w64-mingw32-g++)


# includes
include_directories(
    include
    /usr/x86_64-w64-mingw32/include
)

# Qt
set(CMAKE_PREFIX_PATH "/usr/x86_64-w64-mingw32/include/qt")
find_package(Qt5 COMPONENTS Widgets Quick REQUIRED)

# compile
file(GLOB SOURCES "src/*.cpp")
add_executable(disvoice ${SOURCES})

# link against libraries
target_link_libraries(disvoice /usr/local/lib/libportaudio-2.dll Qt5::Widgets Qt5::Quick)
