package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceDataCmd = &cobra.Command{
	Use:   "finspace-data",
	Short: "The FinSpace APIs let you take actions inside the FinSpace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspaceDataCmd).Standalone()

	})
	rootCmd.AddCommand(finspaceDataCmd)
}
