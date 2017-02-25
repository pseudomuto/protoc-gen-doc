# Building the Plugin

## Prerequisites

* Protocol Buffers library from Google
* QtCore from Qt 5

On Debian/Ubuntu, these packages can be installed with:

    apt install qt5-qmake qt5-default libprotobuf-dev protobuf-compiler libprotoc-dev

## Linux and BSD

At a terminal command prompt, run

    $ qmake
    $ make

in the top-level directory to build the plugin. This will produce the plugin
executable (`protoc-gen-doc`). There's no install step, just copy the executable to
where you want it.

## Windows

Start a Qt/MSVC command prompt, load `vcvarsall.bat` and then run

    > set PROTOBUF_PREFIX=/path/to/protobuf-2.6.1
    > qmake
    > nmake

in the top-level directory to build the plugin. This will produce the plugin
executable (`release\protoc-gen-doc.exe`). `PROTOBUF_PREFIX` is the path to where the
protobuf library was built. You can create a standalone ZIP distribution with `nmake
zip`. MSVC is currently the only supported compiler on Windows. Building with MinGW
should work, but the `zip` target is not available. I'll try to fix this in the
future.

## Mac OS X

### Install Build Tools
First, you need to have Xcode and Xcode Command Line Tools installed. You can follow [this instruction](https://railsapps.github.io/xcode-command-line-tools.html) to install both.

Then, you need to install Homebrew, the mighty package manager on macos, see [Homebrew site](http://brew.sh) for instructions.

Finally, at a Terminal prompt, run:
```
brew update
brew install qt5 protobuf fop docbook-xsl
brew link --force qt5
export PROTOBUF_PREFIX=$(brew --prefix protobuf)
git clone git@github.com:estan/protoc-gen-doc.git
cd protoc-gen-doc
qmake
make
cd examples
make clean
make
```
Basically, this installed packages required to build `protoc-gen-doc`, then clone the master git repo (or you can use your own fork) and build from the source. Finally, it generates the doc from example protobuf files.


In the top-level directory to build the plugin. This will produce the plugin
executable (`protoc-gen-doc`). `PROTOBUF_PREFIX` is the path to where the protobuf
library was installed. There's no install step, just copy the executable to where you
want it, or specify the path to `protoc-gen-doc` with --plugin.

Note that on Mac OS X, the protobuf library should be built with with clang
(`CC=clang` and `CXX=clang++`), or you'll get linker errors.

If you need even more detailed instructions, you can look at the Travis build file (https://github.com/estan/protoc-gen-doc/blob/master/.travis.yml). The tool is built and tested regularly on Mac OS X, and that file contains the exact steps.

