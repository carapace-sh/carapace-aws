package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_listIdMappingJobsCmd = &cobra.Command{
	Use:   "list-id-mapping-jobs",
	Short: "Lists all ID mapping jobs for a given workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_listIdMappingJobsCmd).Standalone()

	entityresolution_listIdMappingJobsCmd.Flags().String("max-results", "", "The maximum number of objects returned per page.")
	entityresolution_listIdMappingJobsCmd.Flags().String("next-token", "", "The pagination token from the previous API call.")
	entityresolution_listIdMappingJobsCmd.Flags().String("workflow-name", "", "The name of the workflow to be retrieved.")
	entityresolution_listIdMappingJobsCmd.MarkFlagRequired("workflow-name")
	entityresolutionCmd.AddCommand(entityresolution_listIdMappingJobsCmd)
}
