package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_listIdMappingWorkflowsCmd = &cobra.Command{
	Use:   "list-id-mapping-workflows",
	Short: "Returns a list of all the `IdMappingWorkflows` that have been created for an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_listIdMappingWorkflowsCmd).Standalone()

	entityresolution_listIdMappingWorkflowsCmd.Flags().String("max-results", "", "The maximum number of objects returned per page.")
	entityresolution_listIdMappingWorkflowsCmd.Flags().String("next-token", "", "The pagination token from the previous API call.")
	entityresolutionCmd.AddCommand(entityresolution_listIdMappingWorkflowsCmd)
}
