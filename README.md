# yh - YAML Highlighter

Dummy YAML syntax highlighter to bring colours where only jq could.

This project starts with the author's incapacity to give up on the absence of colours while playing with kubernetes cli and YAML output (`kubectl get something -o yaml`) as would happen with JSON and jq.

Unable to find a simple YAML highlighter he could throw some content to, with little to none ~~respect~~ expectation, and full of inspiration from his recent Go training, he decided to write one from scratch.

![Comparison with or without yh](https://raw.githubusercontent.com/andreazorzetto/yh/master/images/comparison.png)

# Download & installation

Make your life more colourful from here:

- [MacOS](https://github.com/andreazorzetto/yh/releases/download/v0.2.0/yh-osx-amd64.zip)
- [Linux amd64](https://github.com/andreazorzetto/yh/releases/download/v0.2.0/yh-linux-amd64.zip)
- [Linux 386](https://github.com/andreazorzetto/yh/releases/download/v0.2.0/yh-linux-386.zip)
- [Windows](https://github.com/andreazorzetto/yh/releases/download/v0.2.0/yh-win-amd64.zip)

**Linux or MacOS**

Unzip

`unzip yh-<os>-<arch>.zip`

Move the binary somewhere in your PATH

`mv yh /usr/local/bin/`

**Windows**

Like above but with your mouse

# How to use

As easy as you secretly hoped

`kubectl get pod alpine -o yaml | yh`

Other commands:

- `yh help`
- `yh version`

# Work in progress

This project is still a work in progress. It essentially works but more improvements and docs will follow. 
