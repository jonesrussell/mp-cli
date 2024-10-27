/*
Copyright © 2024 Russell Jones <499552+jonesrussell@users.noreply.github.com>
*/
package main

import (
	"mp-cli/cmd"
	"mp-cli/config"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(config.LoadConfig),
		fx.Invoke(func(cfg *config.Config) {
			cmd.Execute()
		}),
	)
	app.Run()
}
