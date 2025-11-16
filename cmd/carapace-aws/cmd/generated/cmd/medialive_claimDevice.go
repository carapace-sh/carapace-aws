package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_claimDeviceCmd = &cobra.Command{
	Use:   "claim-device",
	Short: "Send a request to claim an AWS Elemental device that you have purchased from a third-party vendor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_claimDeviceCmd).Standalone()

	medialive_claimDeviceCmd.Flags().String("id", "", "The id of the device you want to claim.")
	medialiveCmd.AddCommand(medialive_claimDeviceCmd)
}
