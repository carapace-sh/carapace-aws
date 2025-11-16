package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getSegmentsCmd = &cobra.Command{
	Use:   "get-segments",
	Short: "Retrieves information about the configuration, dimension, and other settings for all the segments that are associated with an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getSegmentsCmd).Standalone()

	pinpoint_getSegmentsCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getSegmentsCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
	pinpoint_getSegmentsCmd.Flags().String("token", "", "The NextToken string that specifies which page of results to return in a paginated response.")
	pinpoint_getSegmentsCmd.MarkFlagRequired("application-id")
	pinpointCmd.AddCommand(pinpoint_getSegmentsCmd)
}
