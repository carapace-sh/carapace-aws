package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_updateLoadBalancerAttributeCmd = &cobra.Command{
	Use:   "update-load-balancer-attribute",
	Short: "Updates the specified attribute for a load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_updateLoadBalancerAttributeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_updateLoadBalancerAttributeCmd).Standalone()

		lightsail_updateLoadBalancerAttributeCmd.Flags().String("attribute-name", "", "The name of the attribute you want to update.")
		lightsail_updateLoadBalancerAttributeCmd.Flags().String("attribute-value", "", "The value that you want to specify for the attribute name.")
		lightsail_updateLoadBalancerAttributeCmd.Flags().String("load-balancer-name", "", "The name of the load balancer that you want to modify (`my-load-balancer`.")
		lightsail_updateLoadBalancerAttributeCmd.MarkFlagRequired("attribute-name")
		lightsail_updateLoadBalancerAttributeCmd.MarkFlagRequired("attribute-value")
		lightsail_updateLoadBalancerAttributeCmd.MarkFlagRequired("load-balancer-name")
	})
	lightsailCmd.AddCommand(lightsail_updateLoadBalancerAttributeCmd)
}
