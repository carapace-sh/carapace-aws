package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createTopicRuleCmd = &cobra.Command{
	Use:   "create-topic-rule",
	Short: "Creates a rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createTopicRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_createTopicRuleCmd).Standalone()

		iot_createTopicRuleCmd.Flags().String("rule-name", "", "The name of the rule.")
		iot_createTopicRuleCmd.Flags().String("tags", "", "Metadata which can be used to manage the topic rule.")
		iot_createTopicRuleCmd.Flags().String("topic-rule-payload", "", "The rule payload.")
		iot_createTopicRuleCmd.MarkFlagRequired("rule-name")
		iot_createTopicRuleCmd.MarkFlagRequired("topic-rule-payload")
	})
	iotCmd.AddCommand(iot_createTopicRuleCmd)
}
