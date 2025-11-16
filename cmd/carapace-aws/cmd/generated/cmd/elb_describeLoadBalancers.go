package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_describeLoadBalancersCmd = &cobra.Command{
	Use:   "describe-load-balancers",
	Short: "Describes the specified the load balancers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_describeLoadBalancersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elb_describeLoadBalancersCmd).Standalone()

		elb_describeLoadBalancersCmd.Flags().String("load-balancer-names", "", "The names of the load balancers.")
		elb_describeLoadBalancersCmd.Flags().String("marker", "", "The marker for the next set of results.")
		elb_describeLoadBalancersCmd.Flags().String("page-size", "", "The maximum number of results to return with this call (a number from 1 to 400).")
	})
	elbCmd.AddCommand(elb_describeLoadBalancersCmd)
}
