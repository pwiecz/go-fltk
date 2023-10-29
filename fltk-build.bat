@ECHO OFF

IF "%GOOS%" == "" (
    ECHO "GOOS environment variable is not set."
    ECHO "This script should not be invoked directly, but run via go generate."
    ECHO "Call: go generate from go-fltk source directory"
    EXIT /B 1
)

IF "%GOARCH%" == "" (
    ECHO "GOARCH environment variable is not set."
    ECHO "This script should not be invoked directly, but run via go generate."
    ECHO "Call: go generate from go-fltk source directory"
    EXIT /B 1
)

ECHO "Building FLTK for OS: %GOOS%, architecture: %GOARCH%"
ECHO "Current dir: %CD%"
SET CMAKE_INSTALL_PREFIX="%CD%"
SET CMAKE_INSTALL_LIBDIR="lib/%GOOS%/%GOARCH%"
SET CMAKE_INSTALL_INCLUDEDIR="include"

IF NOT EXIST fltk_build MKDIR fltk_build
CD fltk_build
IF %ERRORLEVEL% NEQ 0 EXIT /B 1

IF EXIST fltk (
    ECHO "Found existing FLTK directory"
    CD fltk
    IF %ERRORLEVEL% NEQ 0 EXIT /B 1
    git checkout src/Fl_win32.cxx
    IF %ERRORLEVEL% NEQ 0 EXIT /B 1
    git fetch
    IF %ERRORLEVEL% NEQ 0 EXIT /B 1
) ELSE (
    ECHO "Cloning FLTK repository"
    git clone https://github.com/fltk/fltk.git
    IF %ERRORLEVEL% NEQ 0 EXIT /B 1
    CD fltk
    IF %ERRORLEVEL% NEQ 0 EXIT /B 1
)
git apply ../../lib/fltk-1.4.patch
IF %ERRORLEVEL% NEQ 0 EXIT /B 1

SET COMMIT="master"
git checkout "%COMMIT%"
IF %ERRORLEVEL% NEQ 0 EXIT /B 1

CD ..

REM patch -Np1 -i ../lib/fltk-1.4.patch
REM IF %ERRORLEVEL% NEQ 0 EXIT /B 1

cmake -G "MinGW Makefiles" -S fltk -B build -DFLTK_BUILD_TEST=OFF -DCMAKE_BUILD_TYPE=Release -DFLTK_BUILD_EXAMPLES=OFF -DFLTK_BUILD_FLUID=OFF -DFLTK_BUILD_FLTK_OPTIONS=OFF -DOPTION_USE_SYSTEM_LIBJPEG=OFF -DOPTION_USE_SYSTEM_LIBPNG=OFF -DOPTION_USE_SYSTEM_ZLIB=OFF -DCMAKE_INSTALL_PREFIX=%CMAKE_INSTALL_PREFIX% -DCMAKE_INSTALL_INCLUDEDIR=%CMAKE_INSTALL_INCLUDEDIR% -DFLTK_INCLUDEDIR=%CMAKE_INSTALL_PREFIX%/%CMAKE_INSTALL_INCLUDEDIR% -DCMAKE_INSTALL_LIBDIR=%CMAKE_INSTALL_LIBDIR% -DFLTK_LIBDIR=%CMAKE_INSTALL_PREFIX%/%CMAKE_INSTALL_LIBDIR%
IF %ERRORLEVEL% NEQ 0 EXIT /B 1
cmake --build build --parallel
IF %ERRORLEVEL% NEQ 0 EXIT /B 1
cmake --install build
IF %ERRORLEVEL% NEQ 0 EXIT /B 1

CD ..

SET CGO_FILENAME=cgo_%GOOS%_%GOARCH%.go

REM Hardcoding contents of cgo directive for windows, as we cannot extract it from fltk-config if we're not using bash.
ECHO // go:build %GOOS% ^&^& %GOARCH% > %CGO_FILENAME%
ECHO: >>  %CGO_FILENAME%
ECHO package fltk >> %CGO_FILENAME%
ECHO: >> %CGO_FILENAME%
ECHO // #cgo %GOOS%,%GOARCH% CXXFLAGS: -std=c++11 >> %CGO_FILENAME%
ECHO // #cgo %GOOS%,%GOARCH% CPPFLAGS: -I${SRCDIR}/include -I${SRCDIR}/include/FL/images -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64 >> %CGO_FILENAME%
ECHO // #cgo %GOOS%,%GOARCH% LDFLAGS: -mwindows ${SRCDIR}/lib/windows/amd64/libfltk_images.a ${SRCDIR}/lib/windows/amd64/libfltk_jpeg.a ${SRCDIR}/lib/windows/amd64/libfltk_png.a ${SRCDIR}/lib/windows/amd64/libfltk_z.a ${SRCDIR}/lib/windows/amd64/libfltk_gl.a -lglu32 -lopengl32 ${SRCDIR}/lib/windows/amd64/libfltk_forms.a ${SRCDIR}/lib/windows/amd64/libfltk.a -lgdiplus -lole32 -luuid -lcomctl32 -lws2_32 >> %CGO_FILENAME%
ECHO import "C" >> %CGO_FILENAME%

COLOR 07
