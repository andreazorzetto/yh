# yh - YAML Highlighter

A dummy YAML syntax highlighter that brings colours where only jq could.

This project starts with the author's incapacity to give up on colours while playing with kubernetes cli and YAML output (`kubectl get something -o yaml`) as it would happen with JSON and jq.

Unable to find a YAML highlighter he could simply throw some content at, with little to no ~~respect~~ expectation, and full of inspiration from his recent Go training, he decided to write one from scratch.

![Comparison with or without yh](https://raw.githubusercontent.com/andreazorzetto/yh/master/images/comparison.png)

# How to use

As **easy** as you secretly hoped

`kubectl get pod alpine -o yaml | yh`

Other commands:

- `yh help`
- `yh version`

# Download & installation

Make your life more colourful starting from here:

- [MacOS](https://github.com/andreazorzetto/yh/releases/download/v0.2.1/yh-osx-amd64.zip)
- [Linux amd64](https://github.com/andreazorzetto/yh/releases/download/v0.2.1/yh-linux-amd64.zip)
- [Linux 386](https://github.com/andreazorzetto/yh/releases/download/v0.2.1/yh-linux-386.zip)
- [Windows](https://github.com/andreazorzetto/yh/releases/download/v0.2.1/yh-windows-amd64.zip)

**Linux or MacOS**

Unzip

`unzip yh-<os>-<arch>.zip`

Move the binary somewhere in your PATH

`mv yh /usr/local/bin/`

**Windows**

Like above but with your mouse

# Future developments

This project is still a work in progress and I'm considering different routes regarding improvements and new features. Anyone thinking to contribute or make suggestions/requests is more than welcome.
