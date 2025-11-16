package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_describeInputDeviceCmd = &cobra.Command{
	Use:   "describe-input-device",
	Short: "Gets the details for the input device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_describeInputDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_describeInputDeviceCmd).Standalone()

		medialive_describeInputDeviceCmd.Flags().String("input-device-id", "", "The unique ID of this input device.")
		medialive_describeInputDeviceCmd.MarkFlagRequired("input-device-id")
	})
	medialiveCmd.AddCommand(medialive_describeInputDeviceCmd)
}
