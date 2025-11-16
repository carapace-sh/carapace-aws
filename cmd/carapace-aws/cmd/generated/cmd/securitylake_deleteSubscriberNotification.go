package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_deleteSubscriberNotificationCmd = &cobra.Command{
	Use:   "delete-subscriber-notification",
	Short: "Deletes the specified subscription notification in Amazon Security Lake for the organization you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_deleteSubscriberNotificationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securitylake_deleteSubscriberNotificationCmd).Standalone()

		securitylake_deleteSubscriberNotificationCmd.Flags().String("subscriber-id", "", "The ID of the Security Lake subscriber account.")
		securitylake_deleteSubscriberNotificationCmd.MarkFlagRequired("subscriber-id")
	})
	securitylakeCmd.AddCommand(securitylake_deleteSubscriberNotificationCmd)
}
