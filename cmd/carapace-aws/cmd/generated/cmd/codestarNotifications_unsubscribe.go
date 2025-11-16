package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarNotifications_unsubscribeCmd = &cobra.Command{
	Use:   "unsubscribe",
	Short: "Removes an association between a notification rule and an Amazon Q Developer in chat applications topic so that subscribers to that topic stop receiving notifications when the events described in the rule are triggered.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarNotifications_unsubscribeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codestarNotifications_unsubscribeCmd).Standalone()

		codestarNotifications_unsubscribeCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the notification rule.")
		codestarNotifications_unsubscribeCmd.Flags().String("target-address", "", "The ARN of the Amazon Q Developer in chat applications topic to unsubscribe from the notification rule.")
		codestarNotifications_unsubscribeCmd.MarkFlagRequired("arn")
		codestarNotifications_unsubscribeCmd.MarkFlagRequired("target-address")
	})
	codestarNotificationsCmd.AddCommand(codestarNotifications_unsubscribeCmd)
}
