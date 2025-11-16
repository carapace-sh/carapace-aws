package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_disableAvailabilityZonesForLoadBalancerCmd = &cobra.Command{
	Use:   "disable-availability-zones-for-load-balancer",
	Short: "Removes the specified Availability Zones from the set of Availability Zones for the specified load balancer in EC2-Classic or a default VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_disableAvailabilityZonesForLoadBalancerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elb_disableAvailabilityZonesForLoadBalancerCmd).Standalone()

		elb_disableAvailabilityZonesForLoadBalancerCmd.Flags().String("availability-zones", "", "The Availability Zones.")
		elb_disableAvailabilityZonesForLoadBalancerCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
		elb_disableAvailabilityZonesForLoadBalancerCmd.MarkFlagRequired("availability-zones")
		elb_disableAvailabilityZonesForLoadBalancerCmd.MarkFlagRequired("load-balancer-name")
	})
	elbCmd.AddCommand(elb_disableAvailabilityZonesForLoadBalancerCmd)
}
