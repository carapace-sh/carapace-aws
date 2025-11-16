package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shieldCmd = &cobra.Command{
	Use:   "shield",
	Short: "Shield Advanced",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shieldCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(shieldCmd).Standalone()

	})
	rootCmd.AddCommand(shieldCmd)
}
