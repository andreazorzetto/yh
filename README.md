# yh - YAML Highlighter

A dummy syntax highlighter that brings colours to YAML output, jq style.

This project starts with the author's incapacity to give up on colours while playing with kubernetes cli and YAML output (`kubectl get something -o yaml`).

Unable to find another YAML highlighter where he could simply dump something on, with little to no ~~respect~~ expectations, and inspired by the recent Go training, he decided to write one from scratch.

![Comparison with or without yh](https://raw.githubusercontent.com/andreazorzetto/yh/master/images/comparison.png)

# How to use

As **easy** as you secretly hoped

`kubectl get pod alpine -o yaml | yh`

Other commands:

- `yh help`
- `yh version`

# Download & installation

Find the latest releases here:

- [Releases](https://github.com/andreazorzetto/yh/releases)

**Brew**

`brew install yh`

**Linux or MacOS**

Unzip

`unzip yh-<os>-<arch>.zip`

Move the binary somewhere in your PATH

`mv yh /usr/local/bin/`

**Windows**

Like above but with your mouse

# Future developments

The aim of this project is to be a simple highlighting tool, while there are more featured projects our there to do YAML parsing. However I intend to keep maintaining the code and introduce new features and/or explore different routes in the future. Anyone thinking to contribute or make suggestions/requests would be more than welcome.
