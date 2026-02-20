Name:		wlcg-site-snmp
Version:	1.1
Release:	1
Summary:	WLCG Network Traffic Monitoring
BuildArch:	x86_64

URL: https://gitlab.cern.ch/pcollado/wlcg-site-snmp-go

License:	GPLv3

BuildRequires:	systemd

# Longer description on what the package is/does
%description
Implementation of SNMP-based ingress and egress traffic
monitoring for WLCG sites.

It has been written in Go so that it can be easily distributed
as a standalone binary.

# Time to copy the binary file!
%install
# Delete the previos build root
rm -rf %{buildroot}

# Create the necessary directories
mkdir -p %{buildroot}%{_bindir}
mkdir -p %{buildroot}%{_unitdir}
mkdir -p %{buildroot}%{_sysconfdir}/%{name}
mkdir -p %{buildroot}%{_mandir}/man1

# And install the necessary files
install -m 0775 %{_sourcedir}/%{name}         %{buildroot}%{_bindir}/%{name}
install -m 0644 %{_sourcedir}/%{name}.json    %{buildroot}%{_sysconfdir}/%{name}/conf.json
install -m 0664 %{_sourcedir}/%{name}.service %{buildroot}%{_unitdir}/%{name}.service
install -m 0664 %{_sourcedir}/%{name}.1.gz    %{buildroot}%{_mandir}/man1/%{name}.1.gz

# Files provided by the package. Check https://docs.fedoraproject.org/en-US/packaging-guidelines/#_manpages too!
%files
%{_bindir}/%{name}
%config(noreplace) %{_sysconfdir}/%{name}/conf.json
%{_unitdir}/%{name}.service
%{_mandir}/man1/%{name}.1*

# Changes introducd with each version
%changelog
* Sun Oct 22 2023 Pablo Collado Soto <pablo.collado.soto@cern.ch>
- First RPM-packaged version
