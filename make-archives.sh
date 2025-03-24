#!/bin/bash

for p in ./vfs/*/; do 
	name="${p%/}"
	name="${name##*/}"
	output="./vfs/${name}.tar"
	echo "Creating archive $output"
	tar -cf "$output" -C "$p" .
done
