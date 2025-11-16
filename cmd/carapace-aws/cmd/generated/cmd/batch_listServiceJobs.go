package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_listServiceJobsCmd = &cobra.Command{
	Use:   "list-service-jobs",
	Short: "Returns a list of service jobs for a specified job queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_listServiceJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(batch_listServiceJobsCmd).Standalone()

		batch_listServiceJobsCmd.Flags().String("filters", "", "The filter to apply to the query.")
		batch_listServiceJobsCmd.Flags().String("job-queue", "", "The name or ARN of the job queue with which to list service jobs.")
		batch_listServiceJobsCmd.Flags().String("job-status", "", "The job status with which to filter service jobs.")
		batch_listServiceJobsCmd.Flags().String("max-results", "", "The maximum number of results returned by `ListServiceJobs` in paginated output.")
		batch_listServiceJobsCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListServiceJobs` request where `maxResults` was used and the results exceeded the value of that parameter.")
	})
	batchCmd.AddCommand(batch_listServiceJobsCmd)
}
