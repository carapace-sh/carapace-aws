package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_deleteSafetyRuleCmd = &cobra.Command{
	Use:   "delete-safety-rule",
	Short: "Deletes a safety rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_deleteSafetyRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryControlConfig_deleteSafetyRuleCmd).Standalone()

		route53RecoveryControlConfig_deleteSafetyRuleCmd.Flags().String("safety-rule-arn", "", "The ARN of the safety rule.")
		route53RecoveryControlConfig_deleteSafetyRuleCmd.MarkFlagRequired("safety-rule-arn")
	})
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_deleteSafetyRuleCmd)
}
