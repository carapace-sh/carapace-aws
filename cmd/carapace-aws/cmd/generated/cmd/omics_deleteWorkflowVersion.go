package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_deleteWorkflowVersionCmd = &cobra.Command{
	Use:   "delete-workflow-version",
	Short: "Deletes a workflow version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_deleteWorkflowVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_deleteWorkflowVersionCmd).Standalone()

		omics_deleteWorkflowVersionCmd.Flags().String("version-name", "", "The workflow version name.")
		omics_deleteWorkflowVersionCmd.Flags().String("workflow-id", "", "The workflow's ID.")
		omics_deleteWorkflowVersionCmd.MarkFlagRequired("version-name")
		omics_deleteWorkflowVersionCmd.MarkFlagRequired("workflow-id")
	})
	omicsCmd.AddCommand(omics_deleteWorkflowVersionCmd)
}
