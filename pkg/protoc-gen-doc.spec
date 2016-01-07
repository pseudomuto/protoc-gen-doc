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
Version:        0.7
Release:        1%{?dist}
Summary:        Documentation generator plugin for Google Protocol Buffers
License:        BSD-2-Clause
Url:            http://github.com/estan/protoc-gen-doc
Source0:        https://github.com/estan/protoc-gen-doc/archive/v%{version}.tar.gz
BuildRequires:  pkgconfig(Qt5Core)
BuildRequires:  protobuf-devel

%description
Documentation generator plugin for the Google Protocol Buffers compiler
(protoc). The plugin can generate HTML, DocBook or Markdown documentation
from comments in your .proto files.

%prep
%setup -q -n protoc-gen-doc-%{version}

%build
%if 0%{?suse_version}
%qmake5 PREFIX=%{buildroot}/%{_prefix}
%else
qmake-qt5 PREFIX=%{buildroot}/%{_prefix}
%endif
make %{?_smp_mflags}

%install
make install

%files
%defattr(-,root,root)
%{_bindir}/protoc-gen-doc

%changelog
* Thu Jan 7 2016 Elvis Stansvik <elvstone@gmail.com> - 0.7-1
- Update to version 0.7.
* Wed Apr 8 2015 Elvis Stansvik <elvstone@gmail.com> - 0.6-1
- Initial RPM package.
