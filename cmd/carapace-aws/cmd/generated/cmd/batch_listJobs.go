package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_listJobsCmd = &cobra.Command{
	Use:   "list-jobs",
	Short: "Returns a list of Batch jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_listJobsCmd).Standalone()

	batch_listJobsCmd.Flags().String("array-job-id", "", "The job ID for an array job.")
	batch_listJobsCmd.Flags().String("filters", "", "The filter to apply to the query.")
	batch_listJobsCmd.Flags().String("job-queue", "", "The name or full Amazon Resource Name (ARN) of the job queue used to list jobs.")
	batch_listJobsCmd.Flags().String("job-status", "", "The job status used to filter jobs in the specified queue.")
	batch_listJobsCmd.Flags().String("max-results", "", "The maximum number of results returned by `ListJobs` in a paginated output.")
	batch_listJobsCmd.Flags().String("multi-node-job-id", "", "The job ID for a multi-node parallel job.")
	batch_listJobsCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListJobs` request where `maxResults` was used and the results exceeded the value of that parameter.")
	batchCmd.AddCommand(batch_listJobsCmd)
}
