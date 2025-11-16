package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listTopicRulesCmd = &cobra.Command{
	Use:   "list-topic-rules",
	Short: "Lists the rules for the specific topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listTopicRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listTopicRulesCmd).Standalone()

		iot_listTopicRulesCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		iot_listTopicRulesCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
		iot_listTopicRulesCmd.Flags().String("rule-disabled", "", "Specifies whether the rule is disabled.")
		iot_listTopicRulesCmd.Flags().String("topic", "", "The topic.")
	})
	iotCmd.AddCommand(iot_listTopicRulesCmd)
}
