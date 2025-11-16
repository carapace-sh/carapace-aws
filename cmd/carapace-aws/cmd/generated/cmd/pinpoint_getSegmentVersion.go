package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getSegmentVersionCmd = &cobra.Command{
	Use:   "get-segment-version",
	Short: "Retrieves information about the configuration, dimension, and other settings for a specific version of a segment that's associated with an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getSegmentVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getSegmentVersionCmd).Standalone()

		pinpoint_getSegmentVersionCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getSegmentVersionCmd.Flags().String("segment-id", "", "The unique identifier for the segment.")
		pinpoint_getSegmentVersionCmd.Flags().String("version", "", "The unique version number (Version property) for the campaign version.")
		pinpoint_getSegmentVersionCmd.MarkFlagRequired("application-id")
		pinpoint_getSegmentVersionCmd.MarkFlagRequired("segment-id")
		pinpoint_getSegmentVersionCmd.MarkFlagRequired("version")
	})
	pinpointCmd.AddCommand(pinpoint_getSegmentVersionCmd)
}
