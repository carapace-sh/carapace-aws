package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_getWorkflowVersionCmd = &cobra.Command{
	Use:   "get-workflow-version",
	Short: "Gets information about a workflow version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_getWorkflowVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_getWorkflowVersionCmd).Standalone()

		omics_getWorkflowVersionCmd.Flags().String("export", "", "The export format for the workflow.")
		omics_getWorkflowVersionCmd.Flags().String("type", "", "The workflow's type.")
		omics_getWorkflowVersionCmd.Flags().String("version-name", "", "The workflow version name.")
		omics_getWorkflowVersionCmd.Flags().String("workflow-id", "", "The workflow's ID.")
		omics_getWorkflowVersionCmd.Flags().String("workflow-owner-id", "", "The 12-digit account ID of the workflow owner.")
		omics_getWorkflowVersionCmd.MarkFlagRequired("version-name")
		omics_getWorkflowVersionCmd.MarkFlagRequired("workflow-id")
	})
	omicsCmd.AddCommand(omics_getWorkflowVersionCmd)
}
