package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_updateEventRuleCmd = &cobra.Command{
	Use:   "update-event-rule",
	Short: "Updates an existing `EventRule`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_updateEventRuleCmd).Standalone()

	notifications_updateEventRuleCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) to use to update the `EventRule`.")
	notifications_updateEventRuleCmd.Flags().String("event-pattern", "", "An additional event pattern used to further filter the events this `EventRule` receives.")
	notifications_updateEventRuleCmd.Flags().String("regions", "", "A list of Amazon Web Services Regions that sends events to this `EventRule`.")
	notifications_updateEventRuleCmd.MarkFlagRequired("arn")
	notificationsCmd.AddCommand(notifications_updateEventRuleCmd)
}
