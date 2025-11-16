package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteTopicRuleCmd = &cobra.Command{
	Use:   "delete-topic-rule",
	Short: "Deletes the rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteTopicRuleCmd).Standalone()

	iot_deleteTopicRuleCmd.Flags().String("rule-name", "", "The name of the rule.")
	iot_deleteTopicRuleCmd.MarkFlagRequired("rule-name")
	iotCmd.AddCommand(iot_deleteTopicRuleCmd)
}
