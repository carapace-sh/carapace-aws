package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_putInsightRuleCmd = &cobra.Command{
	Use:   "put-insight-rule",
	Short: "Creates a Contributor Insights rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_putInsightRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_putInsightRuleCmd).Standalone()

		cloudwatch_putInsightRuleCmd.Flags().String("apply-on-transformed-logs", "", "Specify `true` to have this rule evaluate log events after they have been transformed by [Log transformation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html).")
		cloudwatch_putInsightRuleCmd.Flags().String("rule-definition", "", "The definition of the rule, as a JSON object.")
		cloudwatch_putInsightRuleCmd.Flags().String("rule-name", "", "A unique name for the rule.")
		cloudwatch_putInsightRuleCmd.Flags().String("rule-state", "", "The state of the rule.")
		cloudwatch_putInsightRuleCmd.Flags().String("tags", "", "A list of key-value pairs to associate with the Contributor Insights rule.")
		cloudwatch_putInsightRuleCmd.MarkFlagRequired("rule-definition")
		cloudwatch_putInsightRuleCmd.MarkFlagRequired("rule-name")
	})
	cloudwatchCmd.AddCommand(cloudwatch_putInsightRuleCmd)
}
