#!/bin/sh

SRC="./cmd"
BUILD="./build"
mkdir -p "${BUILD}"

for dir in "${SRC}"/*/; do
  name=$(basename "${dir}")
  go build -o "${BUILD}/${name}" "${dir}"

  if [ $? -eq 0 ]; then
    echo "go build ${name}"
  else
    echo "error building ${name}."
  fi
done