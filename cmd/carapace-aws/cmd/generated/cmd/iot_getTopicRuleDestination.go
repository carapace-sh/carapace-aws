package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getTopicRuleDestinationCmd = &cobra.Command{
	Use:   "get-topic-rule-destination",
	Short: "Gets information about a topic rule destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getTopicRuleDestinationCmd).Standalone()

	iot_getTopicRuleDestinationCmd.Flags().String("arn", "", "The ARN of the topic rule destination.")
	iot_getTopicRuleDestinationCmd.MarkFlagRequired("arn")
	iotCmd.AddCommand(iot_getTopicRuleDestinationCmd)
}
