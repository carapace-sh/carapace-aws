package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_deleteLoadBalancerCmd = &cobra.Command{
	Use:   "delete-load-balancer",
	Short: "Deletes the specified load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_deleteLoadBalancerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elb_deleteLoadBalancerCmd).Standalone()

		elb_deleteLoadBalancerCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
		elb_deleteLoadBalancerCmd.MarkFlagRequired("load-balancer-name")
	})
	elbCmd.AddCommand(elb_deleteLoadBalancerCmd)
}
