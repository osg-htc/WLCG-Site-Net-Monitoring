# WLCG SNMP-based Site Network Monitoring
This repository contains a Go reimplementation of the original Python script developed
by Alessandra Forti, Shawn Mc Kee et al. on
[github.com/osg-htc/WLCG-Site-Net-Monitoring](https://github.com/osg-htc/WLCG-Site-Net-Monitoring).

First things first, the existing implementation was instrumental to writing this new version.
Thanks a ton for all the guidance!

The driving force behind reimplementing this tool lies on how deploying Python-based programs can
be a bit cumbersome when dealing with dependencies, virtual environments and all those Python quirks.
We would be lying if we didn't accept we are not too Pythonesque as well...

At any rate, this reimplementation aims to be a drop-in replacement for the current script. We
also provide a RPM package whose definition can be found on `wlcg-site-snmp.spec`.

## Installation
Given we provide a RPM package, you can just download it from the repository's
[releases](https://github.com/osg-htc/WLCG-Site-Net-Monitoring/releases) and install it with:

    $ rpm -i wlcg-site-snmp-1.0.0-1.x86_64.rpm

This will also install the `wlcg-site-snmp.service` SystemD unit file so that you can manage the
service through `systemctl(1)`.

## Configuration
The different configuration options are described on `wlcg-site-snmp.1.md`. This specially-formatted
Markdown file is then passed through `pandoc(1)` to generate the `roff(7)` manpage that's shipped
with the RPM. This is achieved with:

    $ pandoc --standalone --to man wlcg-site-snmp.1.md -o wlcg-site-snmp.1

However, the original Markdown file remains very human readable. Once installed, you should be able to
query the manual page by running:

    $ man ./wlcg-site-snmp

Note the leading `./`, otherwise `man` tries to find the manpage on its database...

## The Makefile
This project provides a Makefile automating common tasks such as compiling the program for both
macOS and Linux targets as well as generating the doc. It will also automatically build the RPM
provided the environment has been previously setup through
[`rpmdev-setuptree`](https://rpm-packaging-guide.github.io). Bear in mind the configured RPM
buildroot is defined through the `RPM_BUILDROOT` variable on the Makefile and is, by default,
set to `$HOME/.rpmbuild`.

You can run `make` without arguments to get a list of available targets.

### Compilation
This implementation is written purely in Go. This means compilation is as easy as running `go build`
from this directory. You will, of course, need a working Go toolchain which you can get [here](https://go.dev/dl/).
You can just do that instead of leveraging the `Makefile` in case that's more your jam!

The compilation generates a statically compiled binary that will run on every platform sharing
the architecture and underlying operating system for which the binary was compiled. These two
will default to the architecture and OS of the environment the program was compiled on, but you
can easily alter them. [Open Source](https://opensource.com/article/21/1/go-cross-compiling)
offers a great deal of information on the topic.

## Running
If running the standalone binary, you can just invoke the generated executable passing the path of the
configuration file through the `--conf` flag:

    $ .bin/wlcg-site-snmp --conf /path/to/configuration.json

The program will then attach to the TTY and keep on running until interrupted with `CTRL+C`.

## Thoughts?
Feel free to reach me at <pablo.collado@uam.es> for any errors, and questions!
