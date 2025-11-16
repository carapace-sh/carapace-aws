package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_createSafetyRuleCmd = &cobra.Command{
	Use:   "create-safety-rule",
	Short: "Creates a safety rule in a control panel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_createSafetyRuleCmd).Standalone()

	route53RecoveryControlConfig_createSafetyRuleCmd.Flags().String("assertion-rule", "", "The assertion rule requested.")
	route53RecoveryControlConfig_createSafetyRuleCmd.Flags().String("client-token", "", "A unique, case-sensitive string of up to 64 ASCII characters.")
	route53RecoveryControlConfig_createSafetyRuleCmd.Flags().String("gating-rule", "", "The gating rule requested.")
	route53RecoveryControlConfig_createSafetyRuleCmd.Flags().String("tags", "", "The tags associated with the safety rule.")
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_createSafetyRuleCmd)
}
