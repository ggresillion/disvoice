# Additional clean files
cmake_minimum_required(VERSION 3.16)

if("${CONFIG}" STREQUAL "" OR "${CONFIG}" STREQUAL "Debug")
  file(REMOVE_RECURSE
  "CMakeFiles/disvoice_autogen.dir/AutogenUsed.txt"
  "CMakeFiles/disvoice_autogen.dir/ParseCache.txt"
  "disvoice_autogen"
  )
endif()
