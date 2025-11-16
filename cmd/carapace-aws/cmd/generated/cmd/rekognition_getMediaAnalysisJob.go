package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_getMediaAnalysisJobCmd = &cobra.Command{
	Use:   "get-media-analysis-job",
	Short: "Retrieves the results for a given media analysis job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_getMediaAnalysisJobCmd).Standalone()

	rekognition_getMediaAnalysisJobCmd.Flags().String("job-id", "", "Unique identifier for the media analysis job for which you want to retrieve results.")
	rekognition_getMediaAnalysisJobCmd.MarkFlagRequired("job-id")
	rekognitionCmd.AddCommand(rekognition_getMediaAnalysisJobCmd)
}
