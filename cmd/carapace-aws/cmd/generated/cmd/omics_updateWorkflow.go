package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_updateWorkflowCmd = &cobra.Command{
	Use:   "update-workflow",
	Short: "Updates information about a workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_updateWorkflowCmd).Standalone()

	omics_updateWorkflowCmd.Flags().String("description", "", "A description for the workflow.")
	omics_updateWorkflowCmd.Flags().String("id", "", "The workflow's ID.")
	omics_updateWorkflowCmd.Flags().String("name", "", "A name for the workflow.")
	omics_updateWorkflowCmd.Flags().String("readme-markdown", "", "The markdown content for the workflow's README file.")
	omics_updateWorkflowCmd.Flags().String("storage-capacity", "", "The default static storage capacity (in gibibytes) for runs that use this workflow or workflow version.")
	omics_updateWorkflowCmd.Flags().String("storage-type", "", "The default storage type for runs that use this workflow.")
	omics_updateWorkflowCmd.MarkFlagRequired("id")
	omicsCmd.AddCommand(omics_updateWorkflowCmd)
}
