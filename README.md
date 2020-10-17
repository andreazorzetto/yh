# yh - YAML Highlighter

A dummy syntax highlighter that brings colours to YAML output, jq style.

This project starts with the author's incapacity to give up on colours while playing with kubernetes cli and YAML output (`kubectl get something -o yaml`).

Unable to find another YAML highlighter where to simply dump some code on, with little to no ~~respect~~ expectations, and inspired by my recent Go training, I decided to write one from scratch.

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

**Linux or MacOS**

Unzip

`unzip yh-<os>-<arch>.zip`

Move the binary somewhere in your PATH

`mv yh /usr/local/bin/`

**Windows**

Like above but with your mouse

# Future developments

This project is still a work in progress and I'm considering different routes regarding improvements and new features. Anyone thinking to contribute or make suggestions/requests is more than welcome.
