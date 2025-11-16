package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_deleteWorkflowCmd = &cobra.Command{
	Use:   "delete-workflow",
	Short: "Deletes the specified workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_deleteWorkflowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_deleteWorkflowCmd).Standalone()

		transfer_deleteWorkflowCmd.Flags().String("workflow-id", "", "A unique identifier for the workflow.")
		transfer_deleteWorkflowCmd.MarkFlagRequired("workflow-id")
	})
	transferCmd.AddCommand(transfer_deleteWorkflowCmd)
}
