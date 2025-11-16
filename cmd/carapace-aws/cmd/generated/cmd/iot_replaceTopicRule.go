package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_replaceTopicRuleCmd = &cobra.Command{
	Use:   "replace-topic-rule",
	Short: "Replaces the rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_replaceTopicRuleCmd).Standalone()

	iot_replaceTopicRuleCmd.Flags().String("rule-name", "", "The name of the rule.")
	iot_replaceTopicRuleCmd.Flags().String("topic-rule-payload", "", "The rule payload.")
	iot_replaceTopicRuleCmd.MarkFlagRequired("rule-name")
	iot_replaceTopicRuleCmd.MarkFlagRequired("topic-rule-payload")
	iotCmd.AddCommand(iot_replaceTopicRuleCmd)
}
