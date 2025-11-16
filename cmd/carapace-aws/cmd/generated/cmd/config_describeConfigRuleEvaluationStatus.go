package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeConfigRuleEvaluationStatusCmd = &cobra.Command{
	Use:   "describe-config-rule-evaluation-status",
	Short: "Returns status information for each of your Config managed rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeConfigRuleEvaluationStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_describeConfigRuleEvaluationStatusCmd).Standalone()

		config_describeConfigRuleEvaluationStatusCmd.Flags().String("config-rule-names", "", "The name of the Config managed rules for which you want status information.")
		config_describeConfigRuleEvaluationStatusCmd.Flags().String("limit", "", "The number of rule evaluation results that you want returned.")
		config_describeConfigRuleEvaluationStatusCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
	})
	configCmd.AddCommand(config_describeConfigRuleEvaluationStatusCmd)
}
