# Reverb
[![Man Hours](https://img.shields.io/endpoint?url=https%3A%2F%2Fmh.jessemillar.com%2Fhours%3Frepo%3Dhttps%3A%2F%2Fgithub.com%2Fjessemillar%2Freverb.git)](https://jessemillar.com/r/man-hours) [![Go Report Card](https://goreportcard.com/badge/github.com/jessemillar/reverb)](https://goreportcard.com/report/github.com/jessemillar/reverb)

I used to have a shell function that would allow me to `echo` out messages surrounded by very legible full-terminal-width dividers. This was super useful when I needed to debug a value that would be easily buried in long collections of log output. The shell function worked fine but was difficult to share and use outside my preferred ZSH shell. As a more modern alternative, I built `reverb`. The name `reverb` is a clever guitar-related play on the default `echo` shell command.

## Installation

```
curl -L https://github.com/jessemillar/reverb/releases/latest/download/reverb-linux-amd64 -o reverb && chmod +x reverb && mv reverb /usr/local/bin
```

## Usage

```
> reverb Testing test test
------------------------------------------
Testing test test
------------------------------------------

> reverb
------------------------------------------

> reverb -c ! Warning
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
Warning
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
```
