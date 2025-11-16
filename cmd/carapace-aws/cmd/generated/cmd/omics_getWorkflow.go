package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_getWorkflowCmd = &cobra.Command{
	Use:   "get-workflow",
	Short: "Gets all information about a workflow using its ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_getWorkflowCmd).Standalone()

	omics_getWorkflowCmd.Flags().String("export", "", "The export format for the workflow.")
	omics_getWorkflowCmd.Flags().String("id", "", "The workflow's ID.")
	omics_getWorkflowCmd.Flags().String("type", "", "The workflow's type.")
	omics_getWorkflowCmd.Flags().String("workflow-owner-id", "", "The ID of the workflow owner.")
	omics_getWorkflowCmd.MarkFlagRequired("id")
	omicsCmd.AddCommand(omics_getWorkflowCmd)
}
