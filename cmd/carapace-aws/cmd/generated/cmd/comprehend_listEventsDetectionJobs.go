package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_listEventsDetectionJobsCmd = &cobra.Command{
	Use:   "list-events-detection-jobs",
	Short: "Gets a list of the events detection jobs that you have submitted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_listEventsDetectionJobsCmd).Standalone()

	comprehend_listEventsDetectionJobsCmd.Flags().String("filter", "", "Filters the jobs that are returned.")
	comprehend_listEventsDetectionJobsCmd.Flags().String("max-results", "", "The maximum number of results to return in each page.")
	comprehend_listEventsDetectionJobsCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	comprehendCmd.AddCommand(comprehend_listEventsDetectionJobsCmd)
}
