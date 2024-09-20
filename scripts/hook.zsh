#!/usr/bin/zsh

cmd="$(dirname ${0})/sdkm"

sdkm-export-env() {
	$cmd env go 2>&1 | while IFS='' read -r line; do
		export "${line}" 1>/dev/null 2>&1
	done
}

autoload -U add-zsh-hook
add-zsh-hook chpwd sdkm-export-env

#autoload -U compinit; compinit
#source <($cmd completion zsh)
