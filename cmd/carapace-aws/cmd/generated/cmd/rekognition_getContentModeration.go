package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_getContentModerationCmd = &cobra.Command{
	Use:   "get-content-moderation",
	Short: "Gets the inappropriate, unwanted, or offensive content analysis results for a Amazon Rekognition Video analysis started by [StartContentModeration]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_getContentModerationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_getContentModerationCmd).Standalone()

		rekognition_getContentModerationCmd.Flags().String("aggregate-by", "", "Defines how to aggregate results of the StartContentModeration request.")
		rekognition_getContentModerationCmd.Flags().String("job-id", "", "The identifier for the inappropriate, unwanted, or offensive content moderation job.")
		rekognition_getContentModerationCmd.Flags().String("max-results", "", "Maximum number of results to return per paginated call.")
		rekognition_getContentModerationCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more data to retrieve), Amazon Rekognition returns a pagination token in the response.")
		rekognition_getContentModerationCmd.Flags().String("sort-by", "", "Sort to use for elements in the `ModerationLabelDetections` array.")
		rekognition_getContentModerationCmd.MarkFlagRequired("job-id")
	})
	rekognitionCmd.AddCommand(rekognition_getContentModerationCmd)
}
