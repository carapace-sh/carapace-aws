package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_modifyLoadBalancerAttributesCmd = &cobra.Command{
	Use:   "modify-load-balancer-attributes",
	Short: "Modifies the specified attributes of the specified Application Load Balancer, Network Load Balancer, or Gateway Load Balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_modifyLoadBalancerAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_modifyLoadBalancerAttributesCmd).Standalone()

		elbv2_modifyLoadBalancerAttributesCmd.Flags().String("attributes", "", "The load balancer attributes.")
		elbv2_modifyLoadBalancerAttributesCmd.Flags().String("load-balancer-arn", "", "The Amazon Resource Name (ARN) of the load balancer.")
		elbv2_modifyLoadBalancerAttributesCmd.MarkFlagRequired("attributes")
		elbv2_modifyLoadBalancerAttributesCmd.MarkFlagRequired("load-balancer-arn")
	})
	elbv2Cmd.AddCommand(elbv2_modifyLoadBalancerAttributesCmd)
}
