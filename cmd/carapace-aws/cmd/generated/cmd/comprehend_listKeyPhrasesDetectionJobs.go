package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_listKeyPhrasesDetectionJobsCmd = &cobra.Command{
	Use:   "list-key-phrases-detection-jobs",
	Short: "Get a list of key phrase detection jobs that you have submitted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_listKeyPhrasesDetectionJobsCmd).Standalone()

	comprehend_listKeyPhrasesDetectionJobsCmd.Flags().String("filter", "", "Filters the jobs that are returned.")
	comprehend_listKeyPhrasesDetectionJobsCmd.Flags().String("max-results", "", "The maximum number of results to return in each page.")
	comprehend_listKeyPhrasesDetectionJobsCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	comprehendCmd.AddCommand(comprehend_listKeyPhrasesDetectionJobsCmd)
}
