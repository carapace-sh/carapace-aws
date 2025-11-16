package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarNotifications_subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Creates an association between a notification rule and an Amazon Q Developer in chat applications topic or Amazon Q Developer in chat applications client so that the associated target can receive notifications when the events described in the rule are triggered.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarNotifications_subscribeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codestarNotifications_subscribeCmd).Standalone()

		codestarNotifications_subscribeCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the notification rule for which you want to create the association.")
		codestarNotifications_subscribeCmd.Flags().String("client-request-token", "", "An enumeration token that, when provided in a request, returns the next batch of the results.")
		codestarNotifications_subscribeCmd.Flags().String("target", "", "")
		codestarNotifications_subscribeCmd.MarkFlagRequired("arn")
		codestarNotifications_subscribeCmd.MarkFlagRequired("target")
	})
	codestarNotificationsCmd.AddCommand(codestarNotifications_subscribeCmd)
}
