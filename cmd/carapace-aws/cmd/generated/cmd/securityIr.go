package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIrCmd = &cobra.Command{
	Use:   "security-ir",
	Short: "This guide documents the action and response elements for use of the service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIrCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityIrCmd).Standalone()

	})
	rootCmd.AddCommand(securityIrCmd)
}
