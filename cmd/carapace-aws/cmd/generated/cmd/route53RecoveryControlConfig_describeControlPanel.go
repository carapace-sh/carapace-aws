package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_describeControlPanelCmd = &cobra.Command{
	Use:   "describe-control-panel",
	Short: "Displays details about a control panel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_describeControlPanelCmd).Standalone()

	route53RecoveryControlConfig_describeControlPanelCmd.Flags().String("control-panel-arn", "", "The Amazon Resource Name (ARN) of the control panel.")
	route53RecoveryControlConfig_describeControlPanelCmd.MarkFlagRequired("control-panel-arn")
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_describeControlPanelCmd)
}
