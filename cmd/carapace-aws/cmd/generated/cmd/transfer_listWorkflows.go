package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_listWorkflowsCmd = &cobra.Command{
	Use:   "list-workflows",
	Short: "Lists all workflows associated with your Amazon Web Services account for your current region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_listWorkflowsCmd).Standalone()

	transfer_listWorkflowsCmd.Flags().String("max-results", "", "The maximum number of items to return.")
	transfer_listWorkflowsCmd.Flags().String("next-token", "", "`ListWorkflows` returns the `NextToken` parameter in the output.")
	transferCmd.AddCommand(transfer_listWorkflowsCmd)
}
