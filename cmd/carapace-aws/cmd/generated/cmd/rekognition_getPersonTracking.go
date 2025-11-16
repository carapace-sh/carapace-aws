package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_getPersonTrackingCmd = &cobra.Command{
	Use:   "get-person-tracking",
	Short: "*End of support notice:* On October 31, 2025, AWS will discontinue support for Amazon Rekognition People Pathing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_getPersonTrackingCmd).Standalone()

	rekognition_getPersonTrackingCmd.Flags().String("job-id", "", "The identifier for a job that tracks persons in a video.")
	rekognition_getPersonTrackingCmd.Flags().String("max-results", "", "Maximum number of results to return per paginated call.")
	rekognition_getPersonTrackingCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there are more persons to retrieve), Amazon Rekognition Video returns a pagination token in the response.")
	rekognition_getPersonTrackingCmd.Flags().String("sort-by", "", "Sort to use for elements in the `Persons` array.")
	rekognition_getPersonTrackingCmd.MarkFlagRequired("job-id")
	rekognitionCmd.AddCommand(rekognition_getPersonTrackingCmd)
}
