package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_listRoutingControlsCmd = &cobra.Command{
	Use:   "list-routing-controls",
	Short: "Returns an array of routing controls for a control panel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_listRoutingControlsCmd).Standalone()

	route53RecoveryControlConfig_listRoutingControlsCmd.Flags().String("control-panel-arn", "", "The Amazon Resource Name (ARN) of the control panel.")
	route53RecoveryControlConfig_listRoutingControlsCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
	route53RecoveryControlConfig_listRoutingControlsCmd.Flags().String("next-token", "", "The token that identifies which batch of results you want to see.")
	route53RecoveryControlConfig_listRoutingControlsCmd.MarkFlagRequired("control-panel-arn")
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_listRoutingControlsCmd)
}
