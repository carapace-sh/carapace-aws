package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getSegmentVersionsCmd = &cobra.Command{
	Use:   "get-segment-versions",
	Short: "Retrieves information about the configuration, dimension, and other settings for all the versions of a specific segment that's associated with an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getSegmentVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getSegmentVersionsCmd).Standalone()

		pinpoint_getSegmentVersionsCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getSegmentVersionsCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
		pinpoint_getSegmentVersionsCmd.Flags().String("segment-id", "", "The unique identifier for the segment.")
		pinpoint_getSegmentVersionsCmd.Flags().String("token", "", "The NextToken string that specifies which page of results to return in a paginated response.")
		pinpoint_getSegmentVersionsCmd.MarkFlagRequired("application-id")
		pinpoint_getSegmentVersionsCmd.MarkFlagRequired("segment-id")
	})
	pinpointCmd.AddCommand(pinpoint_getSegmentVersionsCmd)
}
