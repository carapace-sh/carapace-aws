package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_listTopicsDetectionJobsCmd = &cobra.Command{
	Use:   "list-topics-detection-jobs",
	Short: "Gets a list of the topic detection jobs that you have submitted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_listTopicsDetectionJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_listTopicsDetectionJobsCmd).Standalone()

		comprehend_listTopicsDetectionJobsCmd.Flags().String("filter", "", "Filters the jobs that are returned.")
		comprehend_listTopicsDetectionJobsCmd.Flags().String("max-results", "", "The maximum number of results to return in each page.")
		comprehend_listTopicsDetectionJobsCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	})
	comprehendCmd.AddCommand(comprehend_listTopicsDetectionJobsCmd)
}
