package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notificationscontacts_sendActivationCodeCmd = &cobra.Command{
	Use:   "send-activation-code",
	Short: "Sends an activation email to the email address associated with the specified email contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notificationscontacts_sendActivationCodeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notificationscontacts_sendActivationCodeCmd).Standalone()

		notificationscontacts_sendActivationCodeCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the resource.")
		notificationscontacts_sendActivationCodeCmd.MarkFlagRequired("arn")
	})
	notificationscontactsCmd.AddCommand(notificationscontacts_sendActivationCodeCmd)
}
