#!/usr/bin/zsh

cmd="$(dirname ${0})/sdkm"

sdkm-export-env() {
  export $($cmd env go 2>&1) 1>/dev/null 2>&1
}

autoload -U add-zsh-hook
add-zsh-hook chpwd sdkm-export-env

#autoload -U compinit; compinit
#source <($cmd completion zsh)
