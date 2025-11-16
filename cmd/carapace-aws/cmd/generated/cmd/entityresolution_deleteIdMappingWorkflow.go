package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_deleteIdMappingWorkflowCmd = &cobra.Command{
	Use:   "delete-id-mapping-workflow",
	Short: "Deletes the `IdMappingWorkflow` with a given name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_deleteIdMappingWorkflowCmd).Standalone()

	entityresolution_deleteIdMappingWorkflowCmd.Flags().String("workflow-name", "", "The name of the workflow to be deleted.")
	entityresolution_deleteIdMappingWorkflowCmd.MarkFlagRequired("workflow-name")
	entityresolutionCmd.AddCommand(entityresolution_deleteIdMappingWorkflowCmd)
}
