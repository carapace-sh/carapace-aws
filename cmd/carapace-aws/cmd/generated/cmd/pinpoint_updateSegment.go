package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateSegmentCmd = &cobra.Command{
	Use:   "update-segment",
	Short: "Creates a new segment for an application or updates the configuration, dimension, and other settings for an existing segment that's associated with an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateSegmentCmd).Standalone()

	pinpoint_updateSegmentCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_updateSegmentCmd.Flags().String("segment-id", "", "The unique identifier for the segment.")
	pinpoint_updateSegmentCmd.Flags().String("write-segment-request", "", "")
	pinpoint_updateSegmentCmd.MarkFlagRequired("application-id")
	pinpoint_updateSegmentCmd.MarkFlagRequired("segment-id")
	pinpoint_updateSegmentCmd.MarkFlagRequired("write-segment-request")
	pinpointCmd.AddCommand(pinpoint_updateSegmentCmd)
}
