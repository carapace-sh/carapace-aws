package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_updateWorkflowVersionCmd = &cobra.Command{
	Use:   "update-workflow-version",
	Short: "Updates information about the workflow version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_updateWorkflowVersionCmd).Standalone()

	omics_updateWorkflowVersionCmd.Flags().String("description", "", "Description of the workflow version.")
	omics_updateWorkflowVersionCmd.Flags().String("readme-markdown", "", "The markdown content for the workflow version's README file.")
	omics_updateWorkflowVersionCmd.Flags().String("storage-capacity", "", "The default static storage capacity (in gibibytes) for runs that use this workflow version.")
	omics_updateWorkflowVersionCmd.Flags().String("storage-type", "", "The default storage type for runs that use this workflow version.")
	omics_updateWorkflowVersionCmd.Flags().String("version-name", "", "The name of the workflow version.")
	omics_updateWorkflowVersionCmd.Flags().String("workflow-id", "", "The workflow's ID.")
	omics_updateWorkflowVersionCmd.MarkFlagRequired("version-name")
	omics_updateWorkflowVersionCmd.MarkFlagRequired("workflow-id")
	omicsCmd.AddCommand(omics_updateWorkflowVersionCmd)
}
