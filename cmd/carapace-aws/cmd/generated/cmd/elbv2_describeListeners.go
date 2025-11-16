package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_describeListenersCmd = &cobra.Command{
	Use:   "describe-listeners",
	Short: "Describes the specified listeners or the listeners for the specified Application Load Balancer, Network Load Balancer, or Gateway Load Balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_describeListenersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_describeListenersCmd).Standalone()

		elbv2_describeListenersCmd.Flags().String("listener-arns", "", "The Amazon Resource Names (ARN) of the listeners.")
		elbv2_describeListenersCmd.Flags().String("load-balancer-arn", "", "The Amazon Resource Name (ARN) of the load balancer.")
		elbv2_describeListenersCmd.Flags().String("marker", "", "The marker for the next set of results.")
		elbv2_describeListenersCmd.Flags().String("page-size", "", "The maximum number of results to return with this call.")
	})
	elbv2Cmd.AddCommand(elbv2_describeListenersCmd)
}
