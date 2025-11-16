package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteSegmentCmd = &cobra.Command{
	Use:   "delete-segment",
	Short: "Deletes a segment from an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteSegmentCmd).Standalone()

	pinpoint_deleteSegmentCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_deleteSegmentCmd.Flags().String("segment-id", "", "The unique identifier for the segment.")
	pinpoint_deleteSegmentCmd.MarkFlagRequired("application-id")
	pinpoint_deleteSegmentCmd.MarkFlagRequired("segment-id")
	pinpointCmd.AddCommand(pinpoint_deleteSegmentCmd)
}
