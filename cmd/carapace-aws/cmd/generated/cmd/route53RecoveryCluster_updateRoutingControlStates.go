package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryCluster_updateRoutingControlStatesCmd = &cobra.Command{
	Use:   "update-routing-control-states",
	Short: "Set multiple routing control states.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryCluster_updateRoutingControlStatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryCluster_updateRoutingControlStatesCmd).Standalone()

		route53RecoveryCluster_updateRoutingControlStatesCmd.Flags().String("safety-rules-to-override", "", "The Amazon Resource Names (ARNs) for the safety rules that you want to override when you're updating routing control states.")
		route53RecoveryCluster_updateRoutingControlStatesCmd.Flags().String("update-routing-control-state-entries", "", "A set of routing control entries that you want to update.")
		route53RecoveryCluster_updateRoutingControlStatesCmd.MarkFlagRequired("update-routing-control-state-entries")
	})
	route53RecoveryClusterCmd.AddCommand(route53RecoveryCluster_updateRoutingControlStatesCmd)
}
