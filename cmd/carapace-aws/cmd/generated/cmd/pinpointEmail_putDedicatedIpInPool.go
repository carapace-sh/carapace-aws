package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_putDedicatedIpInPoolCmd = &cobra.Command{
	Use:   "put-dedicated-ip-in-pool",
	Short: "Move a dedicated IP address to an existing dedicated IP pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_putDedicatedIpInPoolCmd).Standalone()

	pinpointEmail_putDedicatedIpInPoolCmd.Flags().String("destination-pool-name", "", "The name of the IP pool that you want to add the dedicated IP address to.")
	pinpointEmail_putDedicatedIpInPoolCmd.Flags().String("ip", "", "The IP address that you want to move to the dedicated IP pool.")
	pinpointEmail_putDedicatedIpInPoolCmd.MarkFlagRequired("destination-pool-name")
	pinpointEmail_putDedicatedIpInPoolCmd.MarkFlagRequired("ip")
	pinpointEmailCmd.AddCommand(pinpointEmail_putDedicatedIpInPoolCmd)
}
