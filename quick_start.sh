#!/usr/bin/env bash

set -eoux pipefail


if [ "$(uname)" == "Darwin" ]; then
  brew install graphviz   
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
  sudo apt-get install graphviz
elif [ "$(expr substr $(uname -s) 1 10)" == "MINGW32_NT" ]; then
  echo "Nobody uses 32 bit anymore, you are so outdated... Not worth supporting"
elif [ "$(expr substr $(uname -s) 1 10)" == "MINGW64_NT" ]; then
  choco install graphviz.portable
fi

# Linux
# dep status -dot | dot -T png | display
# MacOS
# dep status -dot | dot -T png | open -f -a /Applications/Preview.app
# Windows
# dep status -dot | dot -T png -o status.png; start status.png