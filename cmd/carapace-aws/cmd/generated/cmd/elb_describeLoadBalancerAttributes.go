package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_describeLoadBalancerAttributesCmd = &cobra.Command{
	Use:   "describe-load-balancer-attributes",
	Short: "Describes the attributes for the specified load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_describeLoadBalancerAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elb_describeLoadBalancerAttributesCmd).Standalone()

		elb_describeLoadBalancerAttributesCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
		elb_describeLoadBalancerAttributesCmd.MarkFlagRequired("load-balancer-name")
	})
	elbCmd.AddCommand(elb_describeLoadBalancerAttributesCmd)
}
