#!/bin/sh

if [ -z "$GOOS" ]; then
    echo "GOOS environment variable is not set."
    echo "This script should not be invoked directly, but run via go generate."
    echo "Call: go generate from go-fltk source directory"
    exit 1
fi
if [ -z "$GOARCH" ]; then
    echo "GOARCH environment variable is not set."
    echo "This script should not be invoked directly, but run via go generate."
    echo "Call: go generate from go-fltk source directory"
    exit 1
fi

echo "Building FLTK for OS: $GOOS, architecture: $GOARCH"

mkdir -p lib/$GOOS/$GOARCH || exit 1

CMAKE_INSTALL_PREFIX="$PWD"
CMAKE_INSTALL_LIBDIR="lib/$GOOS/$GOARCH"
CMAKE_INSTALL_INCLUDEDIR="include"

mkdir -p fltk_build || exit 1
cd fltk_build || exit 1

if [ -d "fltk" ]; then
    echo "Found existing FLTK directory"
    cd fltk || exit 1
    git fetch || exit 1
else
    echo "Cloning FLTK repository"
    git clone https://github.com/fltk/fltk.git || exit 1
    cd fltk || exit 1
fi

#COMMIT="master"
COMMIT="eb759cb118fbf09da51938c04978e609822dbb48"
git checkout "$COMMIT" || exit 1

cd ..

CMAKE_FLAGS="-DFLTK_BUILD_TEST=OFF
             -DCMAKE_BUILD_TYPE=Release
             -DFLTK_BUILD_EXAMPLES=OFF
	     -DFLTK_BUILD_FLUID=OFF
	     -DFLTK_BUILD_FLTK_OPTIONS=OFF
	     -DOPTION_USE_WAYLAND=OFF
             -DOPTION_USE_SYSTEM_LIBJPEG=OFF
             -DOPTION_USE_SYSTEM_LIBPNG=OFF
             -DOPTION_USE_SYSTEM_ZLIB=OFF 
	     -DCMAKE_INSTALL_PREFIX=$CMAKE_INSTALL_PREFIX
	     -DCMAKE_INSTALL_INCLUDEDIR=$CMAKE_INSTALL_INCLUDEDIR
	     -DCMAKE_INSTALL_LIBDIR=$CMAKE_INSTALL_LIBDIR"

case "$GOOS" in 
    "darwin")
	case "$GOARCH" in
	     "amd64")
		 CMAKE_FLAGS="$CMAKE_FLAGS -DCMAKE_OSX_ARCHITECTURES=x86_64"
		 ;;
	     "arm64")
		 CMAKE_FLAGS="$CMAKE_FLAGS -DCMAKE_OSX_ARCHITECTURES=arm64"
		 ;;
	     *)
		 echo "Unsupported MacOS architecture"
		 exit 1
	esac
        ;;
    *)
	;;
esac

cmake -S fltk -B build $CMAKE_FLAGS || exit 1
cmake --build build --parallel || exit 1
cmake --install build || exit 1

cd .. || exit 1

CGO_FILENAME=cgo_${GOOS}_${GOARCH}.go

printf "//go:build $GOOS && $GOARCH\n\npackage fltk\n\n// #cgo $GOOS,$GOARCH CXXFLAGS: -std=c++11\n// #cgo $GOOS,$GOARCH CPPFLAGS: " > $CGO_FILENAME || exit 1

FLTK_CONFIG_FLAGS="--use-gl --use-images --use-forms"

sh fltk_build/build/bin/fltk-config $FLTK_CONFIG_FLAGS --cxxflags | sed s^${CMAKE_INSTALL_PREFIX}^'\$\{SRCDIR\}'^g >> $CGO_FILENAME || exit 1

printf "// #cgo $GOOS,$GOARCH LDFLAGS: " >> $CGO_FILENAME || exit 1

sh fltk_build/build/bin/fltk-config $FLTK_CONFIG_FLAGS --ldstaticflags | sed s^${CMAKE_INSTALL_PREFIX}^'\$\{SRCDIR\}'^g >> $CGO_FILENAME || exit 1

printf 'import "C"\n' >> $CGO_FILENAME || exit 1
