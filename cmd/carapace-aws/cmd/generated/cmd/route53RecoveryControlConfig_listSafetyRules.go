package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_listSafetyRulesCmd = &cobra.Command{
	Use:   "list-safety-rules",
	Short: "List the safety rules (the assertion rules and gating rules) that you've defined for the routing controls in a control panel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_listSafetyRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryControlConfig_listSafetyRulesCmd).Standalone()

		route53RecoveryControlConfig_listSafetyRulesCmd.Flags().String("control-panel-arn", "", "The Amazon Resource Name (ARN) of the control panel.")
		route53RecoveryControlConfig_listSafetyRulesCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
		route53RecoveryControlConfig_listSafetyRulesCmd.Flags().String("next-token", "", "The token that identifies which batch of results you want to see.")
		route53RecoveryControlConfig_listSafetyRulesCmd.MarkFlagRequired("control-panel-arn")
	})
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_listSafetyRulesCmd)
}
