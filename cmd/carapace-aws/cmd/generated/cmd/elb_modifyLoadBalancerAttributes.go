package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_modifyLoadBalancerAttributesCmd = &cobra.Command{
	Use:   "modify-load-balancer-attributes",
	Short: "Modifies the attributes of the specified load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_modifyLoadBalancerAttributesCmd).Standalone()

	elb_modifyLoadBalancerAttributesCmd.Flags().String("load-balancer-attributes", "", "The attributes for the load balancer.")
	elb_modifyLoadBalancerAttributesCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
	elb_modifyLoadBalancerAttributesCmd.MarkFlagRequired("load-balancer-attributes")
	elb_modifyLoadBalancerAttributesCmd.MarkFlagRequired("load-balancer-name")
	elbCmd.AddCommand(elb_modifyLoadBalancerAttributesCmd)
}
