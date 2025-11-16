package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_createWorkflowCmd = &cobra.Command{
	Use:   "create-workflow",
	Short: "Allows you to create a workflow with specified steps and step details the workflow invokes after file transfer completes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_createWorkflowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_createWorkflowCmd).Standalone()

		transfer_createWorkflowCmd.Flags().String("description", "", "A textual description for the workflow.")
		transfer_createWorkflowCmd.Flags().String("on-exception-steps", "", "Specifies the steps (actions) to take if errors are encountered during execution of the workflow.")
		transfer_createWorkflowCmd.Flags().String("steps", "", "Specifies the details for the steps that are in the specified workflow.")
		transfer_createWorkflowCmd.Flags().String("tags", "", "Key-value pairs that can be used to group and search for workflows.")
		transfer_createWorkflowCmd.MarkFlagRequired("steps")
	})
	transferCmd.AddCommand(transfer_createWorkflowCmd)
}
