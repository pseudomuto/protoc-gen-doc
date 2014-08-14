TEMPLATE = app
VERSION = 0.2

CONFIG += debug release console c++11
CONFIG -= app_bundle
QT -= gui

HEADERS += src/mustache.h
SOURCES += src/mustache.cpp src/main.cpp
RESOURCES += protoc-gen-doc.qrc

!win32:QMAKE_CXXFLAGS += -Werror -Wall -Wextra

unix {
    CONFIG += link_pkgconfig
    PKGCONFIG = protobuf
    LIBS += -lprotoc
}

win32-msvc* {
    INCLUDEPATH += "$$(PROTOBUF_PREFIX)\src"

    release:LIBS += "$$(PROTOBUF_PREFIX)\vsprojects\Release\libprotobuf.lib"
    release:LIBS += "$$(PROTOBUF_PREFIX)\vsprojects\Release\libprotoc.lib"

    debug:LIBS += "$$(PROTOBUF_PREFIX)\vsprojects\Debug\libprotobuf.lib"
    debug:LIBS += "$$(PROTOBUF_PREFIX)\vsprojects\Debug\libprotoc.lib"
}

win32-msvc*:release {
    ZIP_DIR = $$shell_path($${TARGET}-$${VERSION}-win32)
    SRC_EXE = $$shell_path(release/$${TARGET}.exe)
    DST_EXE = $$shell_path($${ZIP_DIR}/$${TARGET}.exe)

    win32-msvc2010:MSVCP_DLL = $$shell_path($$(SystemRoot)/SysWOW64/msvcp100.dll)
    win32-msvc2010:MSVCR_DLL = $$shell_path($$(SystemRoot)/SysWOW64/msvcr100.dll)

    win32-msvc2012:MSVCP_DLL = $$shell_path($$(SystemRoot)/SysWOW64/msvcp110.dll)
    win32-msvc2012:MSVCR_DLL = $$shell_path($$(SystemRoot)/SysWOW64/msvcr110.dll)

    win32-msvc2013:MSVCP_DLL = $$shell_path($$(SystemRoot)/SysWOW64/msvcp120.dll)
    win32-msvc2013:MSVCR_DLL = $$shell_path($$(SystemRoot)/SysWOW64/msvcr120.dll)

    WINDEPLOYQT_FLAGS = --release --no-translations --no-compiler-runtime

    zip.target = zip
    zip.depends = zipclean
    zip.commands = md $${ZIP_DIR} && \
                   copy $${SRC_EXE} $${ZIP_DIR} && \
                   copy $${MSVCP_DLL} $${ZIP_DIR} && \
                   copy $${MSVCR_DLL} $${ZIP_DIR} && \
                   windeployqt $${WINDEPLOYQT_FLAGS} $${DST_EXE} && \
                   7z a -r $${ZIP_DIR}.zip $${ZIP_DIR}

    zipclean.target = zipclean
    zipclean.commands = IF EXIST $${ZIP_DIR} rd /s /q $${ZIP_DIR} && \
                        IF EXIST $${ZIP_DIR}.zip del $${ZIP_DIR}.zip

    QMAKE_EXTRA_TARGETS += zip zipclean
}

