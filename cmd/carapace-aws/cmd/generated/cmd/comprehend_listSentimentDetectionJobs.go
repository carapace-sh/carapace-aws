package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_listSentimentDetectionJobsCmd = &cobra.Command{
	Use:   "list-sentiment-detection-jobs",
	Short: "Gets a list of sentiment detection jobs that you have submitted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_listSentimentDetectionJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_listSentimentDetectionJobsCmd).Standalone()

		comprehend_listSentimentDetectionJobsCmd.Flags().String("filter", "", "Filters the jobs that are returned.")
		comprehend_listSentimentDetectionJobsCmd.Flags().String("max-results", "", "The maximum number of results to return in each page.")
		comprehend_listSentimentDetectionJobsCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	})
	comprehendCmd.AddCommand(comprehend_listSentimentDetectionJobsCmd)
}
