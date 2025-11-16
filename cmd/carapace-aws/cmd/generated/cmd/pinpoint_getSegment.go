package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getSegmentCmd = &cobra.Command{
	Use:   "get-segment",
	Short: "Retrieves information about the configuration, dimension, and other settings for a specific segment that's associated with an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getSegmentCmd).Standalone()

	pinpoint_getSegmentCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getSegmentCmd.Flags().String("segment-id", "", "The unique identifier for the segment.")
	pinpoint_getSegmentCmd.MarkFlagRequired("application-id")
	pinpoint_getSegmentCmd.MarkFlagRequired("segment-id")
	pinpointCmd.AddCommand(pinpoint_getSegmentCmd)
}
