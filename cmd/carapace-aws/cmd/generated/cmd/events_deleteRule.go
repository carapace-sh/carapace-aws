package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_deleteRuleCmd = &cobra.Command{
	Use:   "delete-rule",
	Short: "Deletes the specified rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_deleteRuleCmd).Standalone()

	events_deleteRuleCmd.Flags().String("event-bus-name", "", "The name or ARN of the event bus associated with the rule.")
	events_deleteRuleCmd.Flags().Bool("force", false, "If this is a managed rule, created by an Amazon Web Services service on your behalf, you must specify `Force` as `True` to delete the rule.")
	events_deleteRuleCmd.Flags().String("name", "", "The name of the rule.")
	events_deleteRuleCmd.Flags().Bool("no-force", false, "If this is a managed rule, created by an Amazon Web Services service on your behalf, you must specify `Force` as `True` to delete the rule.")
	events_deleteRuleCmd.MarkFlagRequired("name")
	events_deleteRuleCmd.Flag("no-force").Hidden = true
	eventsCmd.AddCommand(events_deleteRuleCmd)
}
