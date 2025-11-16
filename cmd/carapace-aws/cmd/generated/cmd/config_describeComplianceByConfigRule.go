package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeComplianceByConfigRuleCmd = &cobra.Command{
	Use:   "describe-compliance-by-config-rule",
	Short: "Indicates whether the specified Config rules are compliant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeComplianceByConfigRuleCmd).Standalone()

	config_describeComplianceByConfigRuleCmd.Flags().String("compliance-types", "", "Filters the results by compliance.")
	config_describeComplianceByConfigRuleCmd.Flags().String("config-rule-names", "", "Specify one or more Config rule names to filter the results by rule.")
	config_describeComplianceByConfigRuleCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
	configCmd.AddCommand(config_describeComplianceByConfigRuleCmd)
}
