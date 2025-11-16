package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_listCallAnalyticsJobsCmd = &cobra.Command{
	Use:   "list-call-analytics-jobs",
	Short: "Provides a list of Call Analytics jobs that match the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_listCallAnalyticsJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_listCallAnalyticsJobsCmd).Standalone()

		transcribe_listCallAnalyticsJobsCmd.Flags().String("job-name-contains", "", "Returns only the Call Analytics jobs that contain the specified string.")
		transcribe_listCallAnalyticsJobsCmd.Flags().String("max-results", "", "The maximum number of Call Analytics jobs to return in each page of results.")
		transcribe_listCallAnalyticsJobsCmd.Flags().String("next-token", "", "If your `ListCallAnalyticsJobs` request returns more results than can be displayed, `NextToken` is displayed in the response with an associated string.")
		transcribe_listCallAnalyticsJobsCmd.Flags().String("status", "", "Returns only Call Analytics jobs with the specified status.")
	})
	transcribeCmd.AddCommand(transcribe_listCallAnalyticsJobsCmd)
}
