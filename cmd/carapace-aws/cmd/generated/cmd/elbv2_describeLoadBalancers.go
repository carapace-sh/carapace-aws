package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_describeLoadBalancersCmd = &cobra.Command{
	Use:   "describe-load-balancers",
	Short: "Describes the specified load balancers or all of your load balancers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_describeLoadBalancersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_describeLoadBalancersCmd).Standalone()

		elbv2_describeLoadBalancersCmd.Flags().String("load-balancer-arns", "", "The Amazon Resource Names (ARN) of the load balancers.")
		elbv2_describeLoadBalancersCmd.Flags().String("marker", "", "The marker for the next set of results.")
		elbv2_describeLoadBalancersCmd.Flags().String("names", "", "The names of the load balancers.")
		elbv2_describeLoadBalancersCmd.Flags().String("page-size", "", "The maximum number of results to return with this call.")
	})
	elbv2Cmd.AddCommand(elbv2_describeLoadBalancersCmd)
}
