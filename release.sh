#!/bin/bash

OSS=(darwin linux freebsd)
ARCHS=(amd64 386)

mkdir -p src
mkdir -p bin
rm -f bin/1024apis-*

for os in "${OSS[@]}"; do
    for arch in "${ARCHS[@]}"; do
    	echo "Building for $os($arch)"
        GOOS=$os GOARCH=$arch go build
        mv 1024apis bin/1024apis-$os-$arch
    done
done

mv *.go src
mv kube src
mv tmpl src
mv examples src
rm -rf .git .gitignore .wercker.yml release.sh

tar czvf 1024apis.tar.gz bin/*

mv 1024apis.tar.gz $WERCKER_OUTPUT_DIR/1024apis.tar.gz
