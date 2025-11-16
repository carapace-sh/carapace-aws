package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryCluster_listRoutingControlsCmd = &cobra.Command{
	Use:   "list-routing-controls",
	Short: "List routing control names and Amazon Resource Names (ARNs), as well as the routing control state for each routing control, along with the control panel name and control panel ARN for the routing controls.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryCluster_listRoutingControlsCmd).Standalone()

	route53RecoveryCluster_listRoutingControlsCmd.Flags().String("control-panel-arn", "", "The Amazon Resource Name (ARN) of the control panel of the routing controls to list.")
	route53RecoveryCluster_listRoutingControlsCmd.Flags().String("max-results", "", "The number of routing controls objects that you want to return with this call.")
	route53RecoveryCluster_listRoutingControlsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	route53RecoveryClusterCmd.AddCommand(route53RecoveryCluster_listRoutingControlsCmd)
}
