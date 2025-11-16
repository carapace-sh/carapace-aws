package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsmlCmd = &cobra.Command{
	Use:   "cleanroomsml",
	Short: "Welcome to the *Amazon Web Services Clean Rooms ML API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsmlCmd).Standalone()

	rootCmd.AddCommand(cleanroomsmlCmd)
}
