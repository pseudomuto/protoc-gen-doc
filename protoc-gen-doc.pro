TEMPLATE = app
CONFIG += console link_pkgconfig c++11
CONFIG -= app_bundle
QT -= gui

PKGCONFIG = protobuf
QMAKE_LIBS += -lprotoc

!win32 {
  QMAKE_CXXFLAGS += -Werror -Wall -Wextra -Wnon-virtual-dtor
}

HEADERS += src/mustache.h
SOURCES += src/mustache.cpp src/main.cpp
RESOURCES += protoc-gen-doc.qrc
