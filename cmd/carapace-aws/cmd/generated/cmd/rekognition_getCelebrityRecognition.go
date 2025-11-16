package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_getCelebrityRecognitionCmd = &cobra.Command{
	Use:   "get-celebrity-recognition",
	Short: "Gets the celebrity recognition results for a Amazon Rekognition Video analysis started by [StartCelebrityRecognition]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_getCelebrityRecognitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_getCelebrityRecognitionCmd).Standalone()

		rekognition_getCelebrityRecognitionCmd.Flags().String("job-id", "", "Job identifier for the required celebrity recognition analysis.")
		rekognition_getCelebrityRecognitionCmd.Flags().String("max-results", "", "Maximum number of results to return per paginated call.")
		rekognition_getCelebrityRecognitionCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more recognized celebrities to retrieve), Amazon Rekognition Video returns a pagination token in the response.")
		rekognition_getCelebrityRecognitionCmd.Flags().String("sort-by", "", "Sort to use for celebrities returned in `Celebrities` field.")
		rekognition_getCelebrityRecognitionCmd.MarkFlagRequired("job-id")
	})
	rekognitionCmd.AddCommand(rekognition_getCelebrityRecognitionCmd)
}
