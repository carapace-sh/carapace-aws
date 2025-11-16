package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_createEventRuleCmd = &cobra.Command{
	Use:   "create-event-rule",
	Short: "Creates an [`EventRule`](https://docs.aws.amazon.com/notifications/latest/userguide/glossary.html) that is associated with a specified `NotificationConfiguration`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_createEventRuleCmd).Standalone()

	notifications_createEventRuleCmd.Flags().String("event-pattern", "", "An additional event pattern used to further filter the events this `EventRule` receives.")
	notifications_createEventRuleCmd.Flags().String("event-type", "", "The event type to match.")
	notifications_createEventRuleCmd.Flags().String("notification-configuration-arn", "", "The Amazon Resource Name (ARN) of the `NotificationConfiguration` associated with this `EventRule`.")
	notifications_createEventRuleCmd.Flags().String("regions", "", "A list of Amazon Web Services Regions that send events to this `EventRule`.")
	notifications_createEventRuleCmd.Flags().String("source", "", "The matched event source.")
	notifications_createEventRuleCmd.MarkFlagRequired("event-type")
	notifications_createEventRuleCmd.MarkFlagRequired("notification-configuration-arn")
	notifications_createEventRuleCmd.MarkFlagRequired("regions")
	notifications_createEventRuleCmd.MarkFlagRequired("source")
	notificationsCmd.AddCommand(notifications_createEventRuleCmd)
}
