package main

import (
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"github.com/gregleroux/snap-plugin-publisher-warp10/warp10"
)

func main() {
	plugin.StartPublisher(&warp10.Warp10Publisher{}, warp10.Name, warp10.Version)
}
