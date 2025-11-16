package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_deleteEventRuleCmd = &cobra.Command{
	Use:   "delete-event-rule",
	Short: "Deletes an `EventRule`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_deleteEventRuleCmd).Standalone()

	notifications_deleteEventRuleCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the `EventRule` to delete.")
	notifications_deleteEventRuleCmd.MarkFlagRequired("arn")
	notificationsCmd.AddCommand(notifications_deleteEventRuleCmd)
}
