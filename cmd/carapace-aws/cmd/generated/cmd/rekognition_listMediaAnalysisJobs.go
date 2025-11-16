package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_listMediaAnalysisJobsCmd = &cobra.Command{
	Use:   "list-media-analysis-jobs",
	Short: "Returns a list of media analysis jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_listMediaAnalysisJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_listMediaAnalysisJobsCmd).Standalone()

		rekognition_listMediaAnalysisJobsCmd.Flags().String("max-results", "", "The maximum number of results to return per paginated call.")
		rekognition_listMediaAnalysisJobsCmd.Flags().String("next-token", "", "Pagination token, if the previous response was incomplete.")
	})
	rekognitionCmd.AddCommand(rekognition_listMediaAnalysisJobsCmd)
}
