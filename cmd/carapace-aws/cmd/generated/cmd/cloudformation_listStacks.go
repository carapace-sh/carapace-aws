package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listStacksCmd = &cobra.Command{
	Use:   "list-stacks",
	Short: "Returns the summary information for stacks whose status matches the specified `StackStatusFilter`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listStacksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_listStacksCmd).Standalone()

		cloudformation_listStacksCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		cloudformation_listStacksCmd.Flags().String("stack-status-filter", "", "Stack status to use as a filter.")
	})
	cloudformationCmd.AddCommand(cloudformation_listStacksCmd)
}
