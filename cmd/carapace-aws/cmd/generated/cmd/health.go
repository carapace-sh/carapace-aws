package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Health",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(healthCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(healthCmd).Standalone()

	})
	rootCmd.AddCommand(healthCmd)
}
