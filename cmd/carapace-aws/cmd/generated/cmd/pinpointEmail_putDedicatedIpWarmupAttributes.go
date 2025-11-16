package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_putDedicatedIpWarmupAttributesCmd = &cobra.Command{
	Use:   "put-dedicated-ip-warmup-attributes",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_putDedicatedIpWarmupAttributesCmd).Standalone()

	pinpointEmail_putDedicatedIpWarmupAttributesCmd.Flags().String("ip", "", "The dedicated IP address that you want to update the warm-up attributes for.")
	pinpointEmail_putDedicatedIpWarmupAttributesCmd.Flags().String("warmup-percentage", "", "The warm-up percentage that you want to associate with the dedicated IP address.")
	pinpointEmail_putDedicatedIpWarmupAttributesCmd.MarkFlagRequired("ip")
	pinpointEmail_putDedicatedIpWarmupAttributesCmd.MarkFlagRequired("warmup-percentage")
	pinpointEmailCmd.AddCommand(pinpointEmail_putDedicatedIpWarmupAttributesCmd)
}
