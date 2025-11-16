package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_createSegmentCmd = &cobra.Command{
	Use:   "create-segment",
	Short: "Creates a new segment for an application or updates the configuration, dimension, and other settings for an existing segment that's associated with an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_createSegmentCmd).Standalone()

	pinpoint_createSegmentCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_createSegmentCmd.Flags().String("write-segment-request", "", "")
	pinpoint_createSegmentCmd.MarkFlagRequired("application-id")
	pinpoint_createSegmentCmd.MarkFlagRequired("write-segment-request")
	pinpointCmd.AddCommand(pinpoint_createSegmentCmd)
}
