# Building the Plugin

## Prerequisites
* Protocol Buffers library from Google
* QtCore from Qt 5

## Linux and BSD
At a terminal command prompt, run

    $ qmake
    $ make

in the top-level directory to build the plugin. This will produce the
plugin executable (`protoc-gen-doc`). There's no install step, just
copy the executable to where you want it.

## Windows
Start a Qt/MSVC command prompt, load `vcvarsall.bat` and then run

    > set PROTOBUF_PREFIX=/path/to/protobuf-2.6.1
    > qmake
    > nmake

in the top-level directory to build the plugin. This will produce the
plugin executable (`release\protoc-gen-doc.exe`). `PROTOBUF_PREFIX` is
the path to where the protobuf library was built. You can create a
standalone ZIP distribution with `nmake zip`. MSVC is currently the
only supported compiler on Windows. Building with MinGW should work,
but the `zip` target is not available. I'll try to fix this in the
future.

## Mac OS X
At a Terminal prompt, run

    $ export PROTOBUF_PREFIX=/path/to/protobuf-2.6.1
    $ qmake
    $ make

in the top-level directory to build the plugin. This will produce the
plugin executable (`protoc-gen-doc`). `PROTOBUF_PREFIX` is the path to
where the protobuf library was installed. There's no install step,
just copy the executable to where you want it.

Note that on Mac OS X, the protobuf library should be build with with
clang (`CC=clang` and `CXX=clang++`), or you'll get linker errors.

