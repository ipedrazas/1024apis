#!/bin/bash

OSS=(darwin linux freebsd)
ARCHS=(amd64 386)

mkdir -p src
mkdir -p $WERCKER_OUTPUT_DIR/bin
rm -f $WERCKER_OUTPUT_DIR/bin/1024apis-*

for os in "${OSS[@]}"; do
    for arch in "${ARCHS[@]}"; do
    	echo "Building for $os($arch)"
        GOOS=$os GOARCH=$arch go build
        mv 1024apis $WERCKER_OUTPUT_DIR/bin/1024apis-$os-$arch
    done
done

mv *.go src
mv kube src
mv tmpl src
mv examples src
rm -rf .git .gitignore .wercker.yml release.sh
