package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_getDedicatedIpCmd = &cobra.Command{
	Use:   "get-dedicated-ip",
	Short: "Get information about a dedicated IP address, including the name of the dedicated IP pool that it's associated with, as well information about the automatic warm-up process for the address.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_getDedicatedIpCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointEmail_getDedicatedIpCmd).Standalone()

		pinpointEmail_getDedicatedIpCmd.Flags().String("ip", "", "The IP address that you want to obtain more information about.")
		pinpointEmail_getDedicatedIpCmd.MarkFlagRequired("ip")
	})
	pinpointEmailCmd.AddCommand(pinpointEmail_getDedicatedIpCmd)
}
