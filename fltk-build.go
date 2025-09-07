//go:build ignore

package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const commit = "tags/release-1.4.4"

func main() {
	if runtime.GOOS == "" {
		fmt.Println("GOOS environment variable is empty")
		os.Exit(1)
	}
	if runtime.GOARCH == "" {
		fmt.Println("GOARCH environment variable is empty")
		os.Exit(1)
	}
	fmt.Printf("Building FLTK for OS: %s, architecture: %s\n", runtime.GOOS, runtime.GOARCH)

	if _, err := exec.LookPath("git"); err != nil {
		fmt.Printf("Cannot find git binary, %v\n", err)
		os.Exit(1)
	}

	if _, err := exec.LookPath("cmake"); err != nil {
		fmt.Printf("Cannot find cmake binary, %v\n", err)
		os.Exit(1)
	}

	libDir := filepath.Join("lib", runtime.GOOS, runtime.GOARCH)

	if err := os.MkdirAll(libDir, 0750); err != nil {
		fmt.Printf("Could not create directory %s, %v\n", libDir, err)
		os.Exit(1)
	}

	includeDir := filepath.Join("include", runtime.GOOS, runtime.GOARCH)
	if err := os.MkdirAll(includeDir, 0750); err != nil {
		fmt.Printf("Could not create directory %s, %v\n", includeDir, err)
		os.Exit(1)
	}

	if err := os.MkdirAll("fltk_build", 0750); err != nil {
		fmt.Printf("Could not create directory fltk_build, %v\n", err)
		os.Exit(1)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Cannot get current directory, %v\n", err)
		os.Exit(1)
	}

	fltkSourceDir := filepath.Join("fltk_build", "fltk")

	fltkStat, err := os.Stat(fltkSourceDir)
	if errors.Is(err, fs.ErrNotExist) {
		fmt.Println("Cloning FLTK repository")

		cloneCmd := exec.Command("git", "clone", "https://github.com/fltk/fltk.git")
		cloneCmd.Dir = "fltk_build"
		cloneCmd.Stdout = os.Stdout
		cloneCmd.Stderr = os.Stderr
		if err := cloneCmd.Run(); err != nil {
			fmt.Printf("Error cloning FLTK repository, %v\n", err)
			os.Exit(1)
		}
	} else if fltkStat.IsDir() {
		fmt.Println("Found existing FLTK directory")

		if runtime.GOOS == "windows" {
			checkoutCmd := exec.Command("git", "checkout", "src/Fl_win32.cxx")
			checkoutCmd.Dir = fltkSourceDir
			checkoutCmd.Stdout = os.Stdout
			checkoutCmd.Stderr = os.Stderr
			if err := checkoutCmd.Run(); err != nil {
				fmt.Printf("Error checking out Fl_win32.cxx, %v\n", err)
				os.Exit(1)
			}
		}

		fetchCmd := exec.Command("git", "fetch")
		fetchCmd.Dir = fltkSourceDir
		fetchCmd.Stdout = os.Stdout
		fetchCmd.Stderr = os.Stderr
		if err := fetchCmd.Run(); err != nil {
			fmt.Printf("Error fetching FLTK source, %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Location for FLTK source code, %s, is not directory\n", fltkSourceDir)
		os.Exit(1)
	}

	checkoutCmd := exec.Command("git", "checkout", commit)
	checkoutCmd.Dir = fltkSourceDir
	checkoutCmd.Stdout = os.Stdout
	checkoutCmd.Stderr = os.Stderr
	if err := checkoutCmd.Run(); err != nil {
		fmt.Printf("Error checking out FLTK source")
		os.Exit(1)
	}

	if runtime.GOOS == "windows" {
		applyCmd := exec.Command("git", "apply", "../../lib/fltk-1.4.patch")
		applyCmd.Dir = fltkSourceDir
		applyCmd.Stdout = os.Stdout
		applyCmd.Stderr = os.Stderr
		if err := applyCmd.Run(); err != nil {
			fmt.Printf("Error applying patch, %v\n", err)
			os.Exit(1)
		}
	}

	var cmakeGenerator string
	if runtime.GOOS == "windows" {
		cmakeGenerator = "MinGW Makefiles"
	} else {
		cmakeGenerator = "Unix Makefiles"
	}

	cmakeCmd := exec.Command(
		"cmake", "-G", cmakeGenerator, "-S", "fltk", "-B", "build",
		"-DCMAKE_BUILD_TYPE=Release",
		"-DFLTK_BUILD_TEST=OFF",
		"-DFLTK_BUILD_EXAMPLES=OFF",
		"-DFLTK_BUILD_FLUID=OFF",
		"-DFLTK_BUILD_HTML_DOCS=OFF",
		"-DFLTK_BUILD_PDF_DOCS=OFF",
		"-DFLTK_BUILD_FLTK_OPTIONS=OFF",
		"-DFLTK_USE_SYSTEM_LIBJPEG=OFF",
		"-DFLTK_USE_SYSTEM_LIBPNG=OFF",
		"-DFLTK_USE_SYSTEM_ZLIB=OFF",
		"-DCMAKE_INSTALL_PREFIX="+currentDir,
		"-DCMAKE_INSTALL_INCLUDEDIR="+includeDir,
		"-DCMAKE_INSTALL_LIBDIR="+libDir,
		"-DFLTK_INCLUDEDIR="+filepath.Join(currentDir, "include", runtime.GOOS, runtime.GOARCH),
		"-DFLTK_LIBDIR="+filepath.Join(currentDir, "lib", runtime.GOOS, runtime.GOARCH))

	if runtime.GOOS == "darwin" {
		if runtime.GOARCH == "amd64" {
			cmakeCmd.Args = append(cmakeCmd.Args, "-DCMAKE_OSX_ARCHITECTURES=x86_64")
		} else if runtime.GOARCH == "arm64" {
			cmakeCmd.Args = append(cmakeCmd.Args, "-DCMAKE_OSX_ARCHITECTURES=arm64")
		} else {
			fmt.Printf("Unsupported MacOS architecture, %s\n", runtime.GOARCH)
			os.Exit(1)
		}
	}
	cmakeCmd.Dir = "fltk_build"
	cmakeCmd.Stdout = os.Stdout
	cmakeCmd.Stderr = os.Stderr
	if err := cmakeCmd.Run(); err != nil {
		fmt.Printf("Error running cmake, %v\n", err)
		os.Exit(1)
	}

	cmakeBuildArgs := []string{"--build", "build", "--parallel"}
	if runtime.GOOS == "openbsd" {
		cmakeBuildArgs = []string{"--build", "build"}
	}

	cmakeBuildCmd := exec.Command("cmake", cmakeBuildArgs...)
	cmakeBuildCmd.Dir = "fltk_build"
	cmakeBuildCmd.Stdout = os.Stdout
	cmakeBuildCmd.Stderr = os.Stderr
	if err := cmakeBuildCmd.Run(); err != nil {
		fmt.Printf("Error running cmake build, %v\n", err)
		os.Exit(1)
	}

	cmakeInstallCmd := exec.Command("cmake", "--install", "build")
	cmakeInstallCmd.Dir = "fltk_build"
	cmakeInstallCmd.Stdout = os.Stdout
	cmakeInstallCmd.Stderr = os.Stderr
	if err := cmakeInstallCmd.Run(); err != nil {
		fmt.Printf("Error running cmake install, %v\n", err)
		os.Exit(1)
	}

	cgoFilename := "cgo_" + runtime.GOOS + "_" + runtime.GOARCH + ".go"
	cgoFile, err := os.Create(cgoFilename)
	if err != nil {
		fmt.Printf("Error opening cgo file, %s, for writing, %v\n", cgoFilename, err)
		os.Exit(1)
	}
	defer cgoFile.Close()
	fmt.Fprintf(cgoFile, "//go:build %s && %s\n\n", runtime.GOOS, runtime.GOARCH)
	fmt.Fprintln(cgoFile, "package fltk\n")
	fmt.Fprintf(cgoFile, "// #cgo %s,%s CXXFLAGS: -std=c++11\n", runtime.GOOS, runtime.GOARCH)

	if runtime.GOOS != "windows" {
		fltkConfigPath := filepath.Join("fltk_build", "build", "bin", "fltk-config")
		fltkConfigStat, err := os.Stat(fltkConfigPath)
		if err != nil {
			fmt.Printf("Error checking for fltk-config file, %v\n", err)
			os.Exit(1)
		}
		if !fltkConfigStat.Mode().IsRegular() {
			fmt.Println("fltk-config file is not a regular file")
			os.Exit(1)
		}
		if err := os.Chmod(fltkConfigPath, fltkConfigStat.Mode().Perm()|0111); err != nil {
			fmt.Printf("Error making fltk-config executable, %v", err)
			os.Exit(1)
		}
		fltkConfigCxxCmd := exec.Command(fltkConfigPath, "--use-gl", "--use-images", "--use-forms", "--cxxflags")
		cxxOutput, err := fltkConfigCxxCmd.Output()
		if err != nil {
			fmt.Printf("Error running fltk-config --cxxflags, %v\n", err)
			os.Exit(1)
		}
		fltkConfigCxxFlags := strings.ReplaceAll(string(cxxOutput), currentDir, "${SRCDIR}")
		if runtime.GOOS == "openbsd" {
			fltkConfigCxxFlags = "-I/usr/X11R6/include " + fltkConfigCxxFlags
		}
		fmt.Fprintf(cgoFile, "// #cgo %s,%s CPPFLAGS: %s", runtime.GOOS, runtime.GOARCH, fltkConfigCxxFlags)
		if fltkConfigCxxFlags[len(fltkConfigCxxFlags)-1] != '\n' {
			fmt.Fprintln(cgoFile, "")
		}

		fltkConfigLdCmd := exec.Command(fltkConfigPath, "--use-gl", "--use-images", "--use-forms", "--ldstaticflags")
		ldOutput, err := fltkConfigLdCmd.Output()
		if err != nil {
			fmt.Printf("Error running fltk-config --ldstaticflags, %v\n", err)
			os.Exit(1)
		}
		fltkConfigLdFlags := strings.ReplaceAll(string(ldOutput), currentDir, "${SRCDIR}")
		if runtime.GOOS == "openbsd" {
			fltkConfigLdFlags = "-L/usr/X11R6/lib " + fltkConfigLdFlags
		}
		fmt.Fprintf(cgoFile, "// #cgo %s,%s LDFLAGS: %s", runtime.GOOS, runtime.GOARCH, fltkConfigLdFlags)
		if fltkConfigLdFlags[len(fltkConfigLdFlags)-1] != '\n' {
			fmt.Fprintln(cgoFile, "")
		}

	} else {
		// Switching to slashes in paths in cgo directives as backslashes are causing problems when being passed to gcc.
		libDirWithSlashes := filepath.ToSlash(libDir)
		includeDirWithSlashes := filepath.ToSlash(includeDir)
		// Hardcoding contents of cgo directive for windows,
		// as we cannot extract it from fltk-config if we're not using a UNIX shell.
		fmt.Fprintf(cgoFile, "// #cgo %s,%s CPPFLAGS: -I${SRCDIR}/%s -I${SRCDIR}/%s/FL/images -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64\n", runtime.GOOS, runtime.GOARCH, includeDirWithSlashes, includeDirWithSlashes)
		fmt.Fprintf(cgoFile, "// #cgo %s,%s LDFLAGS: -mwindows ${SRCDIR}/%s/libfltk_images.a ${SRCDIR}/%s/libfltk_jpeg.a ${SRCDIR}/%s/libfltk_png.a ${SRCDIR}/%s/libfltk_z.a ${SRCDIR}/%s/libfltk_gl.a -lglu32 -lopengl32 ${SRCDIR}/%s/libfltk_forms.a ${SRCDIR}/%s/libfltk.a -lgdiplus -lole32 -luuid -lcomctl32 -lws2_32 -lwinspool\n", runtime.GOOS, runtime.GOARCH, libDirWithSlashes, libDirWithSlashes, libDirWithSlashes, libDirWithSlashes, libDirWithSlashes, libDirWithSlashes, libDirWithSlashes)
	}
	fmt.Fprintln(cgoFile, "import \"C\"")

	fmt.Printf("Successfully generated libraries for OS: %s, architecture: %s\n", runtime.GOOS, runtime.GOARCH)
}
