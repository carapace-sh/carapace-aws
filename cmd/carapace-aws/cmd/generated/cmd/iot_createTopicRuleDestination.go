package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createTopicRuleDestinationCmd = &cobra.Command{
	Use:   "create-topic-rule-destination",
	Short: "Creates a topic rule destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createTopicRuleDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_createTopicRuleDestinationCmd).Standalone()

		iot_createTopicRuleDestinationCmd.Flags().String("destination-configuration", "", "The topic rule destination configuration.")
		iot_createTopicRuleDestinationCmd.MarkFlagRequired("destination-configuration")
	})
	iotCmd.AddCommand(iot_createTopicRuleDestinationCmd)
}
