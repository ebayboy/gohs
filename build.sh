#!/bin/bash

git submodule foreach  --recursive 'tag="$(git config -f $toplevel/.gitmodules submodule.$name.tag)";[[ -n $tag ]] && git reset --hard  $tag || echo "this module has no tag"'
