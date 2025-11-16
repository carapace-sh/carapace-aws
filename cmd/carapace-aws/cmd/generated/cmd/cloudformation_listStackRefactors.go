package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listStackRefactorsCmd = &cobra.Command{
	Use:   "list-stack-refactors",
	Short: "Lists all account stack refactor operations and their statuses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listStackRefactorsCmd).Standalone()

	cloudformation_listStackRefactorsCmd.Flags().String("execution-status-filter", "", "Execution status to use as a filter.")
	cloudformation_listStackRefactorsCmd.Flags().String("max-results", "", "The maximum number of results to be returned with a single call.")
	cloudformation_listStackRefactorsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	cloudformationCmd.AddCommand(cloudformation_listStackRefactorsCmd)
}
