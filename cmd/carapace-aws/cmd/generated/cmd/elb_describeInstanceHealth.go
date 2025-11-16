package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_describeInstanceHealthCmd = &cobra.Command{
	Use:   "describe-instance-health",
	Short: "Describes the state of the specified instances with respect to the specified load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_describeInstanceHealthCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elb_describeInstanceHealthCmd).Standalone()

		elb_describeInstanceHealthCmd.Flags().String("instances", "", "The IDs of the instances.")
		elb_describeInstanceHealthCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
		elb_describeInstanceHealthCmd.MarkFlagRequired("load-balancer-name")
	})
	elbCmd.AddCommand(elb_describeInstanceHealthCmd)
}
