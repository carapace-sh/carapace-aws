package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_detachInstancesFromLoadBalancerCmd = &cobra.Command{
	Use:   "detach-instances-from-load-balancer",
	Short: "Detaches the specified instances from a Lightsail load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_detachInstancesFromLoadBalancerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_detachInstancesFromLoadBalancerCmd).Standalone()

		lightsail_detachInstancesFromLoadBalancerCmd.Flags().String("instance-names", "", "An array of strings containing the names of the instances you want to detach from the load balancer.")
		lightsail_detachInstancesFromLoadBalancerCmd.Flags().String("load-balancer-name", "", "The name of the Lightsail load balancer.")
		lightsail_detachInstancesFromLoadBalancerCmd.MarkFlagRequired("instance-names")
		lightsail_detachInstancesFromLoadBalancerCmd.MarkFlagRequired("load-balancer-name")
	})
	lightsailCmd.AddCommand(lightsail_detachInstancesFromLoadBalancerCmd)
}
