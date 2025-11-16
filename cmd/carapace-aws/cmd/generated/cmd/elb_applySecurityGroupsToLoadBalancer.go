package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_applySecurityGroupsToLoadBalancerCmd = &cobra.Command{
	Use:   "apply-security-groups-to-load-balancer",
	Short: "Associates one or more security groups with your load balancer in a virtual private cloud (VPC).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_applySecurityGroupsToLoadBalancerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elb_applySecurityGroupsToLoadBalancerCmd).Standalone()

		elb_applySecurityGroupsToLoadBalancerCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
		elb_applySecurityGroupsToLoadBalancerCmd.Flags().String("security-groups", "", "The IDs of the security groups to associate with the load balancer.")
		elb_applySecurityGroupsToLoadBalancerCmd.MarkFlagRequired("load-balancer-name")
		elb_applySecurityGroupsToLoadBalancerCmd.MarkFlagRequired("security-groups")
	})
	elbCmd.AddCommand(elb_applySecurityGroupsToLoadBalancerCmd)
}
