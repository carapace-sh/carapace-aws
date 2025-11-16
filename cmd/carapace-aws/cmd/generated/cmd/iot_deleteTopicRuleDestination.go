package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteTopicRuleDestinationCmd = &cobra.Command{
	Use:   "delete-topic-rule-destination",
	Short: "Deletes a topic rule destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteTopicRuleDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deleteTopicRuleDestinationCmd).Standalone()

		iot_deleteTopicRuleDestinationCmd.Flags().String("arn", "", "The ARN of the topic rule destination to delete.")
		iot_deleteTopicRuleDestinationCmd.MarkFlagRequired("arn")
	})
	iotCmd.AddCommand(iot_deleteTopicRuleDestinationCmd)
}
