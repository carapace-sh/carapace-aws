package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_removeTargetsCmd = &cobra.Command{
	Use:   "remove-targets",
	Short: "Removes the specified targets from the specified rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_removeTargetsCmd).Standalone()

	events_removeTargetsCmd.Flags().String("event-bus-name", "", "The name or ARN of the event bus associated with the rule.")
	events_removeTargetsCmd.Flags().Bool("force", false, "If this is a managed rule, created by an Amazon Web Services service on your behalf, you must specify `Force` as `True` to remove targets.")
	events_removeTargetsCmd.Flags().String("ids", "", "The IDs of the targets to remove from the rule.")
	events_removeTargetsCmd.Flags().Bool("no-force", false, "If this is a managed rule, created by an Amazon Web Services service on your behalf, you must specify `Force` as `True` to remove targets.")
	events_removeTargetsCmd.Flags().String("rule", "", "The name of the rule.")
	events_removeTargetsCmd.MarkFlagRequired("ids")
	events_removeTargetsCmd.Flag("no-force").Hidden = true
	events_removeTargetsCmd.MarkFlagRequired("rule")
	eventsCmd.AddCommand(events_removeTargetsCmd)
}
