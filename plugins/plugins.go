package plugins

import (
	sdkmPluginGo "github.com/dev.itbasis.sdkm/internal/plugins/golang"
	pluginGoConsts "github.com/dev.itbasis.sdkm/internal/plugins/golang/consts"
	sdkmPlugin "github.com/dev.itbasis.sdkm/pkg/plugin"
)

var (
	PluginNames = []string{pluginGoConsts.PluginName}

	Plugins = map[string]sdkmPlugin.GetPluginFunc{
		pluginGoConsts.PluginName: sdkmPluginGo.GetPlugin,
	}
)
