/*
Copyright © 2024 Russell Jones <499552+jonesrussell@users.noreply.github.com>
*/
package main

import (
	"context"
	"mp-cli/cmd"
	"mp-cli/config"
	"os"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(config.LoadConfig),
		fx.Invoke(func(cfg *config.Config) {
			// Execute the command within the fx lifecycle
			cmd.Execute()
		}),
	)

	if err := app.Start(context.Background()); err != nil {
		os.Exit(1)
	}
}
