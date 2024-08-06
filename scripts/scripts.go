package scripts

import "embed"

//go:embed hook.zsh hook.bash go
var scripts embed.FS
