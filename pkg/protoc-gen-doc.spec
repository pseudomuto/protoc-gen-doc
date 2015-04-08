#
# spec file for package protoc-gen-doc
#
# Copyright (c) 2015 Elvis Stansvik <elvstone@gmail.com>
#
# All modifications and additions to the file contributed by third parties
# remain the property of their copyright owners, unless otherwise agreed
# upon. The license for this file, and modifications and additions to the
# file, is the same license as for the pristine package itself (unless the
# license for the pristine package is not an Open Source License, in which
# case the license is the MIT License). An "Open Source License" is a
# license that conforms to the Open Source Definition (Version 1.9)
# published by the Open Source Initiative.

Name:           protoc-gen-doc
Version:        0.5
Release:        1
Summary:        Documentation generator plugin for Google Protocol Buffers
License:        BSD
Url:            http://github.com/estan/protoc-gen-doc
Source0:        https://github.com/estan/protoc-gen-doc/archive/v0.5.tar.gz
BuildRequires:  pkgconfig(Qt5Core)
BuildRequires:  protobuf-devel

%description
Documentation generator plugin for the Google Protocol Buffers compiler
(protoc). The plugin can generate HTML, DocBook or Markdown documentation
from comments in your .proto files.

%prep
%setup -q -n protoc-gen-doc-${version}

%build
%qmake5
make %{?_smp_mflags}

%install
make install PREFIX=%{buildroot}

%files
%defattr(-,root,root)
%{_bindir}/protoc-gen-doc

%changelog
* Wed Apr 8 2015 Elvis Stansvik <elvstone@gmail.com> - 0.5-1
- Initial RPM package.
