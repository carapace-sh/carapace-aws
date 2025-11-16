package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_describeInputDeviceThumbnailCmd = &cobra.Command{
	Use:   "describe-input-device-thumbnail",
	Short: "Get the latest thumbnail data for the input device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_describeInputDeviceThumbnailCmd).Standalone()

	medialive_describeInputDeviceThumbnailCmd.Flags().String("accept", "", "The HTTP Accept header.")
	medialive_describeInputDeviceThumbnailCmd.Flags().String("input-device-id", "", "The unique ID of this input device.")
	medialive_describeInputDeviceThumbnailCmd.MarkFlagRequired("accept")
	medialive_describeInputDeviceThumbnailCmd.MarkFlagRequired("input-device-id")
	medialiveCmd.AddCommand(medialive_describeInputDeviceThumbnailCmd)
}
