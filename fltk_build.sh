#!/bin/sh

git submodule update --init --recursive

cd fltk

platform=""
case "$(uname)" in
    Linux*) platform="linux" ;;
    Darwin*) platform="macos" ;;
    MINGW*|MSYS_NT*|CYGWIN*) platform="windows";;
esac

patch -Np1 -i ../lib/fltk-1.4.patch

CMAKE_FLAGS="-DFLTK_BUILD_TEST=OFF
             -DCMAKE_BUILD_TYPE=Release
             -DFLTK_BUILD_EXAMPLES=OFF
             -DOPTION_USE_SYSTEM_LIBJPEG=OFF
             -DOPTION_USE_SYSTEM_LIBPNG=OFF
             -DOPTION_USE_SYSTEM_ZLIB=OFF "

case "$platform" in 
    "macos")
        CMAKE_FLAGS="$CMAKE_FLAGS -DCMAKE_OSX_ARCHITECTURES=x86_64"
        ;;
    "linux")
        CMAKE_FLAGS="$CMAKE_FLAGS -DOPTION_USE_WAYLAND=OFF -DOPTION_USE_CAIRO=ON -DOPTION_USE_PANGO=ON"
        ;;

esac

cmake -B . $CMAKE_FLAGS
cmake --build . --parallel

