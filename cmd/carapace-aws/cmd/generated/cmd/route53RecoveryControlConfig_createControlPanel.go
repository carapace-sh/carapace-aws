package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_createControlPanelCmd = &cobra.Command{
	Use:   "create-control-panel",
	Short: "Creates a new control panel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_createControlPanelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryControlConfig_createControlPanelCmd).Standalone()

		route53RecoveryControlConfig_createControlPanelCmd.Flags().String("client-token", "", "A unique, case-sensitive string of up to 64 ASCII characters.")
		route53RecoveryControlConfig_createControlPanelCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of the cluster for the control panel.")
		route53RecoveryControlConfig_createControlPanelCmd.Flags().String("control-panel-name", "", "The name of the control panel.")
		route53RecoveryControlConfig_createControlPanelCmd.Flags().String("tags", "", "The tags associated with the control panel.")
		route53RecoveryControlConfig_createControlPanelCmd.MarkFlagRequired("cluster-arn")
		route53RecoveryControlConfig_createControlPanelCmd.MarkFlagRequired("control-panel-name")
	})
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_createControlPanelCmd)
}
