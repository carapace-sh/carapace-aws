package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_describeSafetyRuleCmd = &cobra.Command{
	Use:   "describe-safety-rule",
	Short: "Returns information about a safety rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_describeSafetyRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryControlConfig_describeSafetyRuleCmd).Standalone()

		route53RecoveryControlConfig_describeSafetyRuleCmd.Flags().String("safety-rule-arn", "", "The ARN of the safety rule.")
		route53RecoveryControlConfig_describeSafetyRuleCmd.MarkFlagRequired("safety-rule-arn")
	})
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_describeSafetyRuleCmd)
}
