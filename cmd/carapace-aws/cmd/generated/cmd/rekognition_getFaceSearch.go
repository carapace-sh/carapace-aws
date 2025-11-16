package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_getFaceSearchCmd = &cobra.Command{
	Use:   "get-face-search",
	Short: "Gets the face search results for Amazon Rekognition Video face search started by [StartFaceSearch]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_getFaceSearchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_getFaceSearchCmd).Standalone()

		rekognition_getFaceSearchCmd.Flags().String("job-id", "", "The job identifer for the search request.")
		rekognition_getFaceSearchCmd.Flags().String("max-results", "", "Maximum number of results to return per paginated call.")
		rekognition_getFaceSearchCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more search results to retrieve), Amazon Rekognition Video returns a pagination token in the response.")
		rekognition_getFaceSearchCmd.Flags().String("sort-by", "", "Sort to use for grouping faces in the response.")
		rekognition_getFaceSearchCmd.MarkFlagRequired("job-id")
	})
	rekognitionCmd.AddCommand(rekognition_getFaceSearchCmd)
}
