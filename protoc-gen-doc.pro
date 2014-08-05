TEMPLATE = app
QT -= gui
CONFIG += console link_pkgconfig c++11
CONFIG -= app_bundle

PKGCONFIG = protobuf
QMAKE_LIBS += -lprotoc

!win32 {
  QMAKE_CXXFLAGS += -Werror -Wall -Wextra -Wnon-virtual-dtor
}

DEPENDPATH += src
INCLUDEPATH += src

HEADERS   += src/mustache.h
SOURCES   += src/mustache.cpp src/main.cpp
RESOURCES += templates.qrc
