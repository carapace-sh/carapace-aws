package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_deleteControlPanelCmd = &cobra.Command{
	Use:   "delete-control-panel",
	Short: "Deletes a control panel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_deleteControlPanelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryControlConfig_deleteControlPanelCmd).Standalone()

		route53RecoveryControlConfig_deleteControlPanelCmd.Flags().String("control-panel-arn", "", "The Amazon Resource Name (ARN) of the control panel.")
		route53RecoveryControlConfig_deleteControlPanelCmd.MarkFlagRequired("control-panel-arn")
	})
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_deleteControlPanelCmd)
}
