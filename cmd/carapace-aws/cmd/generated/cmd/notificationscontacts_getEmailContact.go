package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notificationscontacts_getEmailContactCmd = &cobra.Command{
	Use:   "get-email-contact",
	Short: "Returns an email contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notificationscontacts_getEmailContactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notificationscontacts_getEmailContactCmd).Standalone()

		notificationscontacts_getEmailContactCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the email contact to get.")
		notificationscontacts_getEmailContactCmd.MarkFlagRequired("arn")
	})
	notificationscontactsCmd.AddCommand(notificationscontacts_getEmailContactCmd)
}
