package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_listDeviceProfilesCmd = &cobra.Command{
	Use:   "list-device-profiles",
	Short: "Lists the device profiles registered to your AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_listDeviceProfilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_listDeviceProfilesCmd).Standalone()

		iotwireless_listDeviceProfilesCmd.Flags().String("device-profile-type", "", "A filter to list only device profiles that use this type, which can be `LoRaWAN` or `Sidewalk`.")
		iotwireless_listDeviceProfilesCmd.Flags().String("max-results", "", "The maximum number of results to return in this operation.")
		iotwireless_listDeviceProfilesCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
	})
	iotwirelessCmd.AddCommand(iotwireless_listDeviceProfilesCmd)
}
