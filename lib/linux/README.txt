Libraries were built using commands:
For x64:
$ ./configure --enable-localjpeg --enable-localpng --enable-localzlib
For ARM:
$ ./configure --enable-localjpeg --enable-localpng --enable-localzlib --disable-wayland
$ make
