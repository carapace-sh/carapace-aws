package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_setSubnetsCmd = &cobra.Command{
	Use:   "set-subnets",
	Short: "Enables the Availability Zones for the specified public subnets for the specified Application Load Balancer, Network Load Balancer or Gateway Load Balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_setSubnetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_setSubnetsCmd).Standalone()

		elbv2_setSubnetsCmd.Flags().String("enable-prefix-for-ipv6-source-nat", "", "\\[Network Load Balancers with UDP listeners] Indicates whether to use an IPv6 prefix from each subnet for source NAT.")
		elbv2_setSubnetsCmd.Flags().String("ip-address-type", "", "The IP address type.")
		elbv2_setSubnetsCmd.Flags().String("load-balancer-arn", "", "The Amazon Resource Name (ARN) of the load balancer.")
		elbv2_setSubnetsCmd.Flags().String("subnet-mappings", "", "The IDs of the public subnets.")
		elbv2_setSubnetsCmd.Flags().String("subnets", "", "The IDs of the public subnets.")
		elbv2_setSubnetsCmd.MarkFlagRequired("load-balancer-arn")
	})
	elbv2Cmd.AddCommand(elbv2_setSubnetsCmd)
}
