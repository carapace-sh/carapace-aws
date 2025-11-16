package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsCmd = &cobra.Command{
	Use:   "cleanrooms",
	Short: "Welcome to the *Clean Rooms API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsCmd).Standalone()

	rootCmd.AddCommand(cleanroomsCmd)
}
