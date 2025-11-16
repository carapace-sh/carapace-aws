package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryCluster_getRoutingControlStateCmd = &cobra.Command{
	Use:   "get-routing-control-state",
	Short: "Get the state for a routing control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryCluster_getRoutingControlStateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryCluster_getRoutingControlStateCmd).Standalone()

		route53RecoveryCluster_getRoutingControlStateCmd.Flags().String("routing-control-arn", "", "The Amazon Resource Name (ARN) for the routing control that you want to get the state for.")
		route53RecoveryCluster_getRoutingControlStateCmd.MarkFlagRequired("routing-control-arn")
	})
	route53RecoveryClusterCmd.AddCommand(route53RecoveryCluster_getRoutingControlStateCmd)
}
