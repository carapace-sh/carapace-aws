package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_createLoadBalancerCmd = &cobra.Command{
	Use:   "create-load-balancer",
	Short: "Creates an Application Load Balancer, Network Load Balancer, or Gateway Load Balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_createLoadBalancerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_createLoadBalancerCmd).Standalone()

		elbv2_createLoadBalancerCmd.Flags().String("customer-owned-ipv4-pool", "", "\\[Application Load Balancers on Outposts] The ID of the customer-owned address pool (CoIP pool).")
		elbv2_createLoadBalancerCmd.Flags().String("enable-prefix-for-ipv6-source-nat", "", "\\[Network Load Balancers with UDP listeners] Indicates whether to use an IPv6 prefix from each subnet for source NAT.")
		elbv2_createLoadBalancerCmd.Flags().String("ip-address-type", "", "The IP address type.")
		elbv2_createLoadBalancerCmd.Flags().String("ipam-pools", "", "\\[Application Load Balancers] The IPAM pools to use with the load balancer.")
		elbv2_createLoadBalancerCmd.Flags().String("name", "", "The name of the load balancer.")
		elbv2_createLoadBalancerCmd.Flags().String("scheme", "", "The nodes of an Internet-facing load balancer have public IP addresses.")
		elbv2_createLoadBalancerCmd.Flags().String("security-groups", "", "\\[Application Load Balancers and Network Load Balancers] The IDs of the security groups for the load balancer.")
		elbv2_createLoadBalancerCmd.Flags().String("subnet-mappings", "", "The IDs of the subnets.")
		elbv2_createLoadBalancerCmd.Flags().String("subnets", "", "The IDs of the subnets.")
		elbv2_createLoadBalancerCmd.Flags().String("tags", "", "The tags to assign to the load balancer.")
		elbv2_createLoadBalancerCmd.Flags().String("type", "", "The type of load balancer.")
		elbv2_createLoadBalancerCmd.MarkFlagRequired("name")
	})
	elbv2Cmd.AddCommand(elbv2_createLoadBalancerCmd)
}
