package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceCmd = &cobra.Command{
	Use:   "finspace",
	Short: "The FinSpace management service provides the APIs for managing FinSpace environments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspaceCmd).Standalone()

	})
	rootCmd.AddCommand(finspaceCmd)
}
