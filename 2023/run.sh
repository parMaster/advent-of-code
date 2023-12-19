#!/bin/sh

for dir in [0-9]*; do
  if [ -d "$dir" ]; then
    cd "$dir"
	go run ./ "$@"
    cd ..
  fi
done