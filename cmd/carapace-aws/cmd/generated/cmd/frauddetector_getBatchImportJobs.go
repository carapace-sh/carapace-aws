package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_getBatchImportJobsCmd = &cobra.Command{
	Use:   "get-batch-import-jobs",
	Short: "Gets all batch import jobs or a specific job of the specified ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_getBatchImportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_getBatchImportJobsCmd).Standalone()

		frauddetector_getBatchImportJobsCmd.Flags().String("job-id", "", "The ID of the batch import job to get.")
		frauddetector_getBatchImportJobsCmd.Flags().String("max-results", "", "The maximum number of objects to return for request.")
		frauddetector_getBatchImportJobsCmd.Flags().String("next-token", "", "The next token from the previous request.")
	})
	frauddetectorCmd.AddCommand(frauddetector_getBatchImportJobsCmd)
}
