package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_createGraphCmd = &cobra.Command{
	Use:   "create-graph",
	Short: "Creates a new behavior graph for the calling account, and sets that account as the administrator account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_createGraphCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(detective_createGraphCmd).Standalone()

		detective_createGraphCmd.Flags().String("tags", "", "The tags to assign to the new behavior graph.")
	})
	detectiveCmd.AddCommand(detective_createGraphCmd)
}
