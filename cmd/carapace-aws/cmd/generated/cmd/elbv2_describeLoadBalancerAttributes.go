package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_describeLoadBalancerAttributesCmd = &cobra.Command{
	Use:   "describe-load-balancer-attributes",
	Short: "Describes the attributes for the specified Application Load Balancer, Network Load Balancer, or Gateway Load Balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_describeLoadBalancerAttributesCmd).Standalone()

	elbv2_describeLoadBalancerAttributesCmd.Flags().String("load-balancer-arn", "", "The Amazon Resource Name (ARN) of the load balancer.")
	elbv2_describeLoadBalancerAttributesCmd.MarkFlagRequired("load-balancer-arn")
	elbv2Cmd.AddCommand(elbv2_describeLoadBalancerAttributesCmd)
}
