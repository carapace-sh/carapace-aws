package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_listMatchingJobsCmd = &cobra.Command{
	Use:   "list-matching-jobs",
	Short: "Lists all jobs for a given workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_listMatchingJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_listMatchingJobsCmd).Standalone()

		entityresolution_listMatchingJobsCmd.Flags().String("max-results", "", "The maximum number of objects returned per page.")
		entityresolution_listMatchingJobsCmd.Flags().String("next-token", "", "The pagination token from the previous API call.")
		entityresolution_listMatchingJobsCmd.Flags().String("workflow-name", "", "The name of the workflow to be retrieved.")
		entityresolution_listMatchingJobsCmd.MarkFlagRequired("workflow-name")
	})
	entityresolutionCmd.AddCommand(entityresolution_listMatchingJobsCmd)
}
