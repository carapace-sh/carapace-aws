package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_listExecutionsCmd = &cobra.Command{
	Use:   "list-executions",
	Short: "Lists all in-progress executions for the specified workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_listExecutionsCmd).Standalone()

	transfer_listExecutionsCmd.Flags().String("max-results", "", "The maximum number of items to return.")
	transfer_listExecutionsCmd.Flags().String("next-token", "", "`ListExecutions` returns the `NextToken` parameter in the output.")
	transfer_listExecutionsCmd.Flags().String("workflow-id", "", "A unique identifier for the workflow.")
	transfer_listExecutionsCmd.MarkFlagRequired("workflow-id")
	transferCmd.AddCommand(transfer_listExecutionsCmd)
}
