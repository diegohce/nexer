#!/bin/bash

export GOPATH=$(go env GOPATH):$(pwd)

export OLD_PS1="$PS1"
export PS1="(go $(basename $(pwd)))$PS1"

alias deactivate="unset GOPATH; unalias deactivate; export PS1=\"$OLD_PS1\"; unset OLD_PS1"



