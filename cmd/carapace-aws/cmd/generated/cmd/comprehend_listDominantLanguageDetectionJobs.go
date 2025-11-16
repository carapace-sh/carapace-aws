package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_listDominantLanguageDetectionJobsCmd = &cobra.Command{
	Use:   "list-dominant-language-detection-jobs",
	Short: "Gets a list of the dominant language detection jobs that you have submitted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_listDominantLanguageDetectionJobsCmd).Standalone()

	comprehend_listDominantLanguageDetectionJobsCmd.Flags().String("filter", "", "Filters that jobs that are returned.")
	comprehend_listDominantLanguageDetectionJobsCmd.Flags().String("max-results", "", "The maximum number of results to return in each page.")
	comprehend_listDominantLanguageDetectionJobsCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	comprehendCmd.AddCommand(comprehend_listDominantLanguageDetectionJobsCmd)
}
