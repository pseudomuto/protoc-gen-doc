# VS_V is the NNN-format Visual Studio version number.
win32-msvc2010:VS_V = 100
win32-msvc2012:VS_V = 110
win32-msvc2013:VS_V = 120

# Find the Visual Studio redistributable folder.
VS_TOOLS = $$getenv(VS$${VS_V}COMNTOOLS)
VS_REDIST = $${VS_TOOLS}/../../VC/redist/x86/Microsoft.VC$${VS_V}.CRT

!exists($${VS_REDIST}) {
    error("Could not find Visual C++ redistributable directory!")
}

# List of files to bundle.
FILES += "release/$${TARGET}.exe"
FILES += "$${VS_REDIST}/msvcr$${VS_V}.dll"
FILES += "$${VS_REDIST}/msvcp$${VS_V}.dll"
FILES += "$$[QT_INSTALL_BINS]/Qt5Core.dll"
FILES += "$$[QT_INSTALL_BINS]/icu*.dll"

# Settings for the zip target.
ZIP_DIR = $${TARGET}-$${VERSION}-win32
ZIP_FILE = $${ZIP_DIR}.zip

# TARGET: zip
#
#  1. Installs FILES to ZIP_DIR
#  2. Compresses all .exe and .dll files with UPX
#  3. Compresses ZIP_DIR into ZIP_DIR.zip with 7zip
#
zip.target = zip
zip.depends = first zipclean
zip.commands = mkdir $${ZIP_DIR} $${ZIP_DIR}\platforms
for (FILE, FILES) {
    zip.commands += && copy $$shell_quote($$shell_path($${FILE})) $${ZIP_DIR}
}
zip.commands += && upx --mono -q $${ZIP_DIR}\*.exe $${ZIP_DIR}\*.dll
zip.commands += && 7z a -r $${ZIP_FILE} $${ZIP_DIR}

# TARGET: cleanzip
#
# Removes ZIP_DIR and ZIP_DIR.zip, if they exist.
#
zipclean.target = zipclean
zipclean.commands = IF EXIST $${ZIP_DIR} rmdir /s /q $${ZIP_DIR} && \
                    IF EXIST $${ZIP_FILE} del $${ZIP_FILE}

QMAKE_EXTRA_TARGETS += zip zipclean
