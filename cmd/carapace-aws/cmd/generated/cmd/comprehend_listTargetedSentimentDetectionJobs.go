package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_listTargetedSentimentDetectionJobsCmd = &cobra.Command{
	Use:   "list-targeted-sentiment-detection-jobs",
	Short: "Gets a list of targeted sentiment detection jobs that you have submitted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_listTargetedSentimentDetectionJobsCmd).Standalone()

	comprehend_listTargetedSentimentDetectionJobsCmd.Flags().String("filter", "", "Filters the jobs that are returned.")
	comprehend_listTargetedSentimentDetectionJobsCmd.Flags().String("max-results", "", "The maximum number of results to return in each page.")
	comprehend_listTargetedSentimentDetectionJobsCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	comprehendCmd.AddCommand(comprehend_listTargetedSentimentDetectionJobsCmd)
}
