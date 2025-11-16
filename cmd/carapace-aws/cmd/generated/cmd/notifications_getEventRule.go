package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_getEventRuleCmd = &cobra.Command{
	Use:   "get-event-rule",
	Short: "Returns a specified `EventRule`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_getEventRuleCmd).Standalone()

	notifications_getEventRuleCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the `EventRule` to return.")
	notifications_getEventRuleCmd.MarkFlagRequired("arn")
	notificationsCmd.AddCommand(notifications_getEventRuleCmd)
}
