TEMPLATE = app
VERSION = 0.2

CONFIG += console c++11
CONFIG -= app_bundle
QT -= gui

HEADERS += src/mustache.h
SOURCES += src/mustache.cpp src/main.cpp
RESOURCES += protoc-gen-doc.qrc

unix {
    # Use pkg-config to find libprotobuf.
    CONFIG += link_pkgconfig
    PKGCONFIG = protobuf

    LIBS += -lprotoc # Has no .pc, so add manually.
}

msvc {
    # Get location of protobuf/protoc libraries.
    PROTOBUF_PREFIX = $$getenv(PROTOBUF_PREFIX)
    isEmpty(PROTOBUF_PREFIX) {
        error(You must set the PROTOBUF_PREFIX environment variable!)
    }

    # Add protobuf/protoc paths to INCLUDEPATH and LIBS.
    INCLUDEPATH += "$${PROTOBUF_PREFIX}\src"
    release:LIBS += "$${PROTOBUF_PREFIX}\vsprojects\Release\libprotobuf.lib"
    release:LIBS += "$${PROTOBUF_PREFIX}\vsprojects\Release\libprotoc.lib"
    debug:LIBS += "$${PROTOBUF_PREFIX}\vsprojects\Debug\libprotobuf.lib"
    debug:LIBS += "$${PROTOBUF_PREFIX}\vsprojects\Debug\libprotoc.lib"

    # Maintain Windows XP compatibility on Visual Studio 2012 and higher.
    QMAKE_LFLAGS += /SUBSYSTEM:CONSOLE,5.01

    # Add zip target in release mode.
    release:include(protoc-gen-doc-win32-zip.pri)
}


# Increase warnings G++ / clang warnings.
*g++*|*clang*:QMAKE_CXXFLAGS += -Werror -Wall -Wextra
