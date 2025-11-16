package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_setIpAddressTypeCmd = &cobra.Command{
	Use:   "set-ip-address-type",
	Short: "Sets the type of IP addresses used by the subnets of the specified load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_setIpAddressTypeCmd).Standalone()

	elbv2_setIpAddressTypeCmd.Flags().String("ip-address-type", "", "The IP address type.")
	elbv2_setIpAddressTypeCmd.Flags().String("load-balancer-arn", "", "The Amazon Resource Name (ARN) of the load balancer.")
	elbv2_setIpAddressTypeCmd.MarkFlagRequired("ip-address-type")
	elbv2_setIpAddressTypeCmd.MarkFlagRequired("load-balancer-arn")
	elbv2Cmd.AddCommand(elbv2_setIpAddressTypeCmd)
}
