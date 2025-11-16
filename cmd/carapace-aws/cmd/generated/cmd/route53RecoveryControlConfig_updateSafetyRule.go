package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_updateSafetyRuleCmd = &cobra.Command{
	Use:   "update-safety-rule",
	Short: "Update a safety rule (an assertion rule or gating rule).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_updateSafetyRuleCmd).Standalone()

	route53RecoveryControlConfig_updateSafetyRuleCmd.Flags().String("assertion-rule-update", "", "The assertion rule to update.")
	route53RecoveryControlConfig_updateSafetyRuleCmd.Flags().String("gating-rule-update", "", "The gating rule to update.")
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_updateSafetyRuleCmd)
}
