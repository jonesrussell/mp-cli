/*
Copyright © 2024 Russell Jones <499552+jonesrussell@users.noreply.github.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewHelloCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "hello",
		Short: "Prints Hello, world!",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, world!")
		},
	}
}
