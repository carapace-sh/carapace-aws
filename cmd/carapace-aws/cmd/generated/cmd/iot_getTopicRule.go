package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getTopicRuleCmd = &cobra.Command{
	Use:   "get-topic-rule",
	Short: "Gets information about the rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getTopicRuleCmd).Standalone()

	iot_getTopicRuleCmd.Flags().String("rule-name", "", "The name of the rule.")
	iot_getTopicRuleCmd.MarkFlagRequired("rule-name")
	iotCmd.AddCommand(iot_getTopicRuleCmd)
}
