package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_confirmTopicRuleDestinationCmd = &cobra.Command{
	Use:   "confirm-topic-rule-destination",
	Short: "Confirms a topic rule destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_confirmTopicRuleDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_confirmTopicRuleDestinationCmd).Standalone()

		iot_confirmTopicRuleDestinationCmd.Flags().String("confirmation-token", "", "The token used to confirm ownership or access to the topic rule confirmation URL.")
		iot_confirmTopicRuleDestinationCmd.MarkFlagRequired("confirmation-token")
	})
	iotCmd.AddCommand(iot_confirmTopicRuleDestinationCmd)
}
