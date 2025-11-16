package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_getIdMappingWorkflowCmd = &cobra.Command{
	Use:   "get-id-mapping-workflow",
	Short: "Returns the `IdMappingWorkflow` with a given name, if it exists.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_getIdMappingWorkflowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_getIdMappingWorkflowCmd).Standalone()

		entityresolution_getIdMappingWorkflowCmd.Flags().String("workflow-name", "", "The name of the workflow.")
		entityresolution_getIdMappingWorkflowCmd.MarkFlagRequired("workflow-name")
	})
	entityresolutionCmd.AddCommand(entityresolution_getIdMappingWorkflowCmd)
}
