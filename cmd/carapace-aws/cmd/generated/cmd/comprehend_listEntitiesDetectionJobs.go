package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_listEntitiesDetectionJobsCmd = &cobra.Command{
	Use:   "list-entities-detection-jobs",
	Short: "Gets a list of the entity detection jobs that you have submitted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_listEntitiesDetectionJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_listEntitiesDetectionJobsCmd).Standalone()

		comprehend_listEntitiesDetectionJobsCmd.Flags().String("filter", "", "Filters the jobs that are returned.")
		comprehend_listEntitiesDetectionJobsCmd.Flags().String("max-results", "", "The maximum number of results to return in each page.")
		comprehend_listEntitiesDetectionJobsCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	})
	comprehendCmd.AddCommand(comprehend_listEntitiesDetectionJobsCmd)
}
