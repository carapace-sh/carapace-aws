package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_disableTopicRuleCmd = &cobra.Command{
	Use:   "disable-topic-rule",
	Short: "Disables the rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_disableTopicRuleCmd).Standalone()

	iot_disableTopicRuleCmd.Flags().String("rule-name", "", "The name of the rule to disable.")
	iot_disableTopicRuleCmd.MarkFlagRequired("rule-name")
	iotCmd.AddCommand(iot_disableTopicRuleCmd)
}
