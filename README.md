# yh - YAML Highlighter

Dummy YAML syntax highlighter to bring colours where only jq could.

A bit of history. This project starts with the author's incapacity to give up on seeing colours while playing with kubernetes command line and YAML output (kubectl get something -o yaml) as would instead happen using JSON format.

Unable to find a simple YAML highlighter where he could just pipe some content with little to none expectations, he decided to write one as an excuse for some Go training.

# Download & installation

Make your life more colourful from here:

- [MacOS](https://github.com/andreazorzetto/yh/releases/download/v0.1/yh-osx-amd64.zip)

- [Linux amd64](https://github.com/andreazorzetto/yh/releases/download/v0.1/yh-linux-amd64.zip)

- [Linux 386](https://github.com/andreazorzetto/yh/releases/download/v0.1/yh-linux-386.zip)

- [Windows](https://github.com/andreazorzetto/yh/releases/download/v0.1/yh-win-amd64.zip)

Unzip `yh-osx-amd64.zip` and copy the binary in `/usr/local/bin` or somewhere else

# Run it

### echo "yaml juice" | yh

Example: 

`kubectl get pod alpine -o yaml | yh`

# Work in progress

This project is still a work in progress. Essentialy works but more improvements and docs will follow
