# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.19

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Disable VCS-based implicit rules.
% : %,v


# Disable VCS-based implicit rules.
% : RCS/%


# Disable VCS-based implicit rules.
% : RCS/%,v


# Disable VCS-based implicit rules.
% : SCCS/s.%


# Disable VCS-based implicit rules.
% : s.%


.SUFFIXES: .hpux_make_needs_suffix_list


# Command-line flag to silence nested $(MAKE).
$(VERBOSE)MAKESILENT = -s

#Suppress display of executed commands.
$(VERBOSE).SILENT:

# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/bin/cmake

# The command to remove a file.
RM = /usr/bin/cmake -E rm -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/guillaume/Projects/disvoice

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/guillaume/Projects/disvoice/build

# Include any dependencies generated for this target.
include CMakeFiles/disvoice.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/disvoice.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/disvoice.dir/flags.make

CMakeFiles/disvoice.dir/src/audio.cpp.o: CMakeFiles/disvoice.dir/flags.make
CMakeFiles/disvoice.dir/src/audio.cpp.o: ../src/audio.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/guillaume/Projects/disvoice/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object CMakeFiles/disvoice.dir/src/audio.cpp.o"
	x86_64-w64-mingw32-g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/disvoice.dir/src/audio.cpp.o -c /home/guillaume/Projects/disvoice/src/audio.cpp

CMakeFiles/disvoice.dir/src/audio.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/disvoice.dir/src/audio.cpp.i"
	x86_64-w64-mingw32-g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/guillaume/Projects/disvoice/src/audio.cpp > CMakeFiles/disvoice.dir/src/audio.cpp.i

CMakeFiles/disvoice.dir/src/audio.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/disvoice.dir/src/audio.cpp.s"
	x86_64-w64-mingw32-g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/guillaume/Projects/disvoice/src/audio.cpp -o CMakeFiles/disvoice.dir/src/audio.cpp.s

CMakeFiles/disvoice.dir/src/common.cpp.o: CMakeFiles/disvoice.dir/flags.make
CMakeFiles/disvoice.dir/src/common.cpp.o: ../src/common.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/guillaume/Projects/disvoice/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Building CXX object CMakeFiles/disvoice.dir/src/common.cpp.o"
	x86_64-w64-mingw32-g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/disvoice.dir/src/common.cpp.o -c /home/guillaume/Projects/disvoice/src/common.cpp

CMakeFiles/disvoice.dir/src/common.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/disvoice.dir/src/common.cpp.i"
	x86_64-w64-mingw32-g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/guillaume/Projects/disvoice/src/common.cpp > CMakeFiles/disvoice.dir/src/common.cpp.i

CMakeFiles/disvoice.dir/src/common.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/disvoice.dir/src/common.cpp.s"
	x86_64-w64-mingw32-g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/guillaume/Projects/disvoice/src/common.cpp -o CMakeFiles/disvoice.dir/src/common.cpp.s

CMakeFiles/disvoice.dir/src/main.cpp.o: CMakeFiles/disvoice.dir/flags.make
CMakeFiles/disvoice.dir/src/main.cpp.o: ../src/main.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/guillaume/Projects/disvoice/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_3) "Building CXX object CMakeFiles/disvoice.dir/src/main.cpp.o"
	x86_64-w64-mingw32-g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/disvoice.dir/src/main.cpp.o -c /home/guillaume/Projects/disvoice/src/main.cpp

CMakeFiles/disvoice.dir/src/main.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/disvoice.dir/src/main.cpp.i"
	x86_64-w64-mingw32-g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/guillaume/Projects/disvoice/src/main.cpp > CMakeFiles/disvoice.dir/src/main.cpp.i

CMakeFiles/disvoice.dir/src/main.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/disvoice.dir/src/main.cpp.s"
	x86_64-w64-mingw32-g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/guillaume/Projects/disvoice/src/main.cpp -o CMakeFiles/disvoice.dir/src/main.cpp.s

CMakeFiles/disvoice.dir/src/vst_plugin.cpp.o: CMakeFiles/disvoice.dir/flags.make
CMakeFiles/disvoice.dir/src/vst_plugin.cpp.o: ../src/vst_plugin.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/guillaume/Projects/disvoice/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_4) "Building CXX object CMakeFiles/disvoice.dir/src/vst_plugin.cpp.o"
	x86_64-w64-mingw32-g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/disvoice.dir/src/vst_plugin.cpp.o -c /home/guillaume/Projects/disvoice/src/vst_plugin.cpp

CMakeFiles/disvoice.dir/src/vst_plugin.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/disvoice.dir/src/vst_plugin.cpp.i"
	x86_64-w64-mingw32-g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/guillaume/Projects/disvoice/src/vst_plugin.cpp > CMakeFiles/disvoice.dir/src/vst_plugin.cpp.i

CMakeFiles/disvoice.dir/src/vst_plugin.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/disvoice.dir/src/vst_plugin.cpp.s"
	x86_64-w64-mingw32-g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/guillaume/Projects/disvoice/src/vst_plugin.cpp -o CMakeFiles/disvoice.dir/src/vst_plugin.cpp.s

CMakeFiles/disvoice.dir/src/window.cpp.o: CMakeFiles/disvoice.dir/flags.make
CMakeFiles/disvoice.dir/src/window.cpp.o: ../src/window.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/guillaume/Projects/disvoice/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_5) "Building CXX object CMakeFiles/disvoice.dir/src/window.cpp.o"
	x86_64-w64-mingw32-g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/disvoice.dir/src/window.cpp.o -c /home/guillaume/Projects/disvoice/src/window.cpp

CMakeFiles/disvoice.dir/src/window.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/disvoice.dir/src/window.cpp.i"
	x86_64-w64-mingw32-g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/guillaume/Projects/disvoice/src/window.cpp > CMakeFiles/disvoice.dir/src/window.cpp.i

CMakeFiles/disvoice.dir/src/window.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/disvoice.dir/src/window.cpp.s"
	x86_64-w64-mingw32-g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/guillaume/Projects/disvoice/src/window.cpp -o CMakeFiles/disvoice.dir/src/window.cpp.s

# Object files for target disvoice
disvoice_OBJECTS = \
"CMakeFiles/disvoice.dir/src/audio.cpp.o" \
"CMakeFiles/disvoice.dir/src/common.cpp.o" \
"CMakeFiles/disvoice.dir/src/main.cpp.o" \
"CMakeFiles/disvoice.dir/src/vst_plugin.cpp.o" \
"CMakeFiles/disvoice.dir/src/window.cpp.o"

# External object files for target disvoice
disvoice_EXTERNAL_OBJECTS =

disvoice: CMakeFiles/disvoice.dir/src/audio.cpp.o
disvoice: CMakeFiles/disvoice.dir/src/common.cpp.o
disvoice: CMakeFiles/disvoice.dir/src/main.cpp.o
disvoice: CMakeFiles/disvoice.dir/src/vst_plugin.cpp.o
disvoice: CMakeFiles/disvoice.dir/src/window.cpp.o
disvoice: CMakeFiles/disvoice.dir/build.make
disvoice: /usr/local/lib/libportaudio-2.dll
disvoice: CMakeFiles/disvoice.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/guillaume/Projects/disvoice/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_6) "Linking CXX executable disvoice"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/disvoice.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/disvoice.dir/build: disvoice

.PHONY : CMakeFiles/disvoice.dir/build

CMakeFiles/disvoice.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/disvoice.dir/cmake_clean.cmake
.PHONY : CMakeFiles/disvoice.dir/clean

CMakeFiles/disvoice.dir/depend:
	cd /home/guillaume/Projects/disvoice/build && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/guillaume/Projects/disvoice /home/guillaume/Projects/disvoice /home/guillaume/Projects/disvoice/build /home/guillaume/Projects/disvoice/build /home/guillaume/Projects/disvoice/build/CMakeFiles/disvoice.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/disvoice.dir/depend

