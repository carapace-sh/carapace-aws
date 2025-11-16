package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_updateControlPanelCmd = &cobra.Command{
	Use:   "update-control-panel",
	Short: "Updates a control panel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_updateControlPanelCmd).Standalone()

	route53RecoveryControlConfig_updateControlPanelCmd.Flags().String("control-panel-arn", "", "The Amazon Resource Name (ARN) of the control panel.")
	route53RecoveryControlConfig_updateControlPanelCmd.Flags().String("control-panel-name", "", "The name of the control panel.")
	route53RecoveryControlConfig_updateControlPanelCmd.MarkFlagRequired("control-panel-arn")
	route53RecoveryControlConfig_updateControlPanelCmd.MarkFlagRequired("control-panel-name")
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_updateControlPanelCmd)
}
