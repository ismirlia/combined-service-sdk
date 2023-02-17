#!/bin/bash

NAMES=($(ls | grep "\-sp\-"))

for name in ${NAMES[@]}; do
    mv $name ${name//\-sp\-/\-ppc\-} 
done
