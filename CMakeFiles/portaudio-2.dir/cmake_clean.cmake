file(REMOVE_RECURSE
  "libportaudio-2.pdb"
  "libportaudio-2.so"
)

# Per-language clean rules from dependency scanning.
foreach(lang )
  include(CMakeFiles/portaudio-2.dir/cmake_clean_${lang}.cmake OPTIONAL)
endforeach()
