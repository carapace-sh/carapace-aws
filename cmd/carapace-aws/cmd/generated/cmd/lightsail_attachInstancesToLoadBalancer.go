package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_attachInstancesToLoadBalancerCmd = &cobra.Command{
	Use:   "attach-instances-to-load-balancer",
	Short: "Attaches one or more Lightsail instances to a load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_attachInstancesToLoadBalancerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_attachInstancesToLoadBalancerCmd).Standalone()

		lightsail_attachInstancesToLoadBalancerCmd.Flags().String("instance-names", "", "An array of strings representing the instance name(s) you want to attach to your load balancer.")
		lightsail_attachInstancesToLoadBalancerCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
		lightsail_attachInstancesToLoadBalancerCmd.MarkFlagRequired("instance-names")
		lightsail_attachInstancesToLoadBalancerCmd.MarkFlagRequired("load-balancer-name")
	})
	lightsailCmd.AddCommand(lightsail_attachInstancesToLoadBalancerCmd)
}
