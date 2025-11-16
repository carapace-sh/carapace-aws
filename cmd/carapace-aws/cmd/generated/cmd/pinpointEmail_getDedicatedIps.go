package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_getDedicatedIpsCmd = &cobra.Command{
	Use:   "get-dedicated-ips",
	Short: "List the dedicated IP addresses that are associated with your Amazon Pinpoint account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_getDedicatedIpsCmd).Standalone()

	pinpointEmail_getDedicatedIpsCmd.Flags().String("next-token", "", "A token returned from a previous call to `GetDedicatedIps` to indicate the position of the dedicated IP pool in the list of IP pools.")
	pinpointEmail_getDedicatedIpsCmd.Flags().String("page-size", "", "The number of results to show in a single call to `GetDedicatedIpsRequest`.")
	pinpointEmail_getDedicatedIpsCmd.Flags().String("pool-name", "", "The name of the IP pool that the dedicated IP address is associated with.")
	pinpointEmailCmd.AddCommand(pinpointEmail_getDedicatedIpsCmd)
}
