package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateTopicRuleDestinationCmd = &cobra.Command{
	Use:   "update-topic-rule-destination",
	Short: "Updates a topic rule destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateTopicRuleDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_updateTopicRuleDestinationCmd).Standalone()

		iot_updateTopicRuleDestinationCmd.Flags().String("arn", "", "The ARN of the topic rule destination.")
		iot_updateTopicRuleDestinationCmd.Flags().String("status", "", "The status of the topic rule destination.")
		iot_updateTopicRuleDestinationCmd.MarkFlagRequired("arn")
		iot_updateTopicRuleDestinationCmd.MarkFlagRequired("status")
	})
	iotCmd.AddCommand(iot_updateTopicRuleDestinationCmd)
}
