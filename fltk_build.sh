#!/bin/sh

commit="master"
fltk_url="https://github.com/fltk/fltk/archive/refs/heads/$commit.tar.gz"

wget -O - "$fltk_url" | tar -xz

platform=""
arch="$(uname -i)"

case "$(uname)" in
    Linux*) 
        platform="linux" 
        ;;
    Darwin*) 
        platform="macos" 
        ;;
    MINGW*|MSYS_NT*|CYGWIN*) 
        platform="windows"
        ;;
esac

CMAKE_FLAGS="-DFLTK_BUILD_TEST=OFF
             -DCMAKE_BUILD_TYPE=Release
             -DFLTK_BUILD_EXAMPLES=OFF
             -DOPTION_USE_SYSTEM_LIBJPEG=OFF
             -DOPTION_USE_SYSTEM_LIBPNG=OFF
             -DOPTION_USE_SYSTEM_ZLIB=OFF "

case "$platform" in 
    "macos")
        CMAKE_FLAGS="$CMAKE_FLAGS -DCMAKE_OSX_ARCHITECTURES=$arch"
        ;;
    "linux")
        CMAKE_FLAGS="$CMAKE_FLAGS -DOPTION_USE_WAYLAND=OFF -DOPTION_USE_CAIRO=ON -DOPTION_USE_PANGO=ON"
        ;;
    "windows")
        patch -Np1 -i ../lib/fltk-1.4.patch
        ;;
esac

cmake -S "fltk-$commit" -B fltk_bin $CMAKE_FLAGS
cmake --build fltk_bin --parallel

