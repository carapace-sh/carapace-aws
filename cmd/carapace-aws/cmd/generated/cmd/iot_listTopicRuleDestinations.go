package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listTopicRuleDestinationsCmd = &cobra.Command{
	Use:   "list-topic-rule-destinations",
	Short: "Lists all the topic rule destinations in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listTopicRuleDestinationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listTopicRuleDestinationsCmd).Standalone()

		iot_listTopicRuleDestinationsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iot_listTopicRuleDestinationsCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
	})
	iotCmd.AddCommand(iot_listTopicRuleDestinationsCmd)
}
