package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_deleteLoadBalancerCmd = &cobra.Command{
	Use:   "delete-load-balancer",
	Short: "Deletes the specified Application Load Balancer, Network Load Balancer, or Gateway Load Balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_deleteLoadBalancerCmd).Standalone()

	elbv2_deleteLoadBalancerCmd.Flags().String("load-balancer-arn", "", "The Amazon Resource Name (ARN) of the load balancer.")
	elbv2_deleteLoadBalancerCmd.MarkFlagRequired("load-balancer-arn")
	elbv2Cmd.AddCommand(elbv2_deleteLoadBalancerCmd)
}
