package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getComplianceSummaryByConfigRuleCmd = &cobra.Command{
	Use:   "get-compliance-summary-by-config-rule",
	Short: "Returns the number of Config rules that are compliant and noncompliant, up to a maximum of 25 for each.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getComplianceSummaryByConfigRuleCmd).Standalone()

	configCmd.AddCommand(config_getComplianceSummaryByConfigRuleCmd)
}
