package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notificationscontacts_deleteEmailContactCmd = &cobra.Command{
	Use:   "delete-email-contact",
	Short: "Deletes an email contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notificationscontacts_deleteEmailContactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notificationscontacts_deleteEmailContactCmd).Standalone()

		notificationscontacts_deleteEmailContactCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the resource.")
		notificationscontacts_deleteEmailContactCmd.MarkFlagRequired("arn")
	})
	notificationscontactsCmd.AddCommand(notificationscontacts_deleteEmailContactCmd)
}
