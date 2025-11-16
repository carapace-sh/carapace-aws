package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_enableAvailabilityZonesForLoadBalancerCmd = &cobra.Command{
	Use:   "enable-availability-zones-for-load-balancer",
	Short: "Adds the specified Availability Zones to the set of Availability Zones for the specified load balancer in EC2-Classic or a default VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_enableAvailabilityZonesForLoadBalancerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elb_enableAvailabilityZonesForLoadBalancerCmd).Standalone()

		elb_enableAvailabilityZonesForLoadBalancerCmd.Flags().String("availability-zones", "", "The Availability Zones.")
		elb_enableAvailabilityZonesForLoadBalancerCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
		elb_enableAvailabilityZonesForLoadBalancerCmd.MarkFlagRequired("availability-zones")
		elb_enableAvailabilityZonesForLoadBalancerCmd.MarkFlagRequired("load-balancer-name")
	})
	elbCmd.AddCommand(elb_enableAvailabilityZonesForLoadBalancerCmd)
}
