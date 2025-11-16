package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_enableTopicRuleCmd = &cobra.Command{
	Use:   "enable-topic-rule",
	Short: "Enables the rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_enableTopicRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_enableTopicRuleCmd).Standalone()

		iot_enableTopicRuleCmd.Flags().String("rule-name", "", "The name of the topic rule to enable.")
		iot_enableTopicRuleCmd.MarkFlagRequired("rule-name")
	})
	iotCmd.AddCommand(iot_enableTopicRuleCmd)
}
