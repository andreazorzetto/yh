# yh - YAML Highlighter

Dummy YAML syntax highlighter to bring colours where only jq could.

A bit of history.

This project starts with the author's incapacity to give up on seeing colours while playing with kubernetes (kubectl) command line and YAML output (-o yaml) as would instead happen using JSON output.

Unable to find a simple YAML highlighter where he could just pipe some content with little to none expectations, he decided to write one as an excuse for some Go training.

# Download

Make your life more colourful from here:

(MacOS only for now) https://github.com/andreazorzetto/yh/releases

# Run it

Example: 
kubectl get pod alpine -o yaml | yh

# Work in progress

This project is still a work in progress and more improvements and docs will follow
