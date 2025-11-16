package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryCluster_updateRoutingControlStateCmd = &cobra.Command{
	Use:   "update-routing-control-state",
	Short: "Set the state of the routing control to reroute traffic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryCluster_updateRoutingControlStateCmd).Standalone()

	route53RecoveryCluster_updateRoutingControlStateCmd.Flags().String("routing-control-arn", "", "The Amazon Resource Name (ARN) for the routing control that you want to update the state for.")
	route53RecoveryCluster_updateRoutingControlStateCmd.Flags().String("routing-control-state", "", "The state of the routing control.")
	route53RecoveryCluster_updateRoutingControlStateCmd.Flags().String("safety-rules-to-override", "", "The Amazon Resource Names (ARNs) for the safety rules that you want to override when you're updating the state of a routing control.")
	route53RecoveryCluster_updateRoutingControlStateCmd.MarkFlagRequired("routing-control-arn")
	route53RecoveryCluster_updateRoutingControlStateCmd.MarkFlagRequired("routing-control-state")
	route53RecoveryClusterCmd.AddCommand(route53RecoveryCluster_updateRoutingControlStateCmd)
}
