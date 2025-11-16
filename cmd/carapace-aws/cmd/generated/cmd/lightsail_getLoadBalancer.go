package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getLoadBalancerCmd = &cobra.Command{
	Use:   "get-load-balancer",
	Short: "Returns information about the specified Lightsail load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getLoadBalancerCmd).Standalone()

	lightsail_getLoadBalancerCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
	lightsail_getLoadBalancerCmd.MarkFlagRequired("load-balancer-name")
	lightsailCmd.AddCommand(lightsail_getLoadBalancerCmd)
}
