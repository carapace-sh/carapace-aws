package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notificationscontacts_createEmailContactCmd = &cobra.Command{
	Use:   "create-email-contact",
	Short: "Creates an email contact for the provided email address.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notificationscontacts_createEmailContactCmd).Standalone()

	notificationscontacts_createEmailContactCmd.Flags().String("email-address", "", "The email address this email contact points to.")
	notificationscontacts_createEmailContactCmd.Flags().String("name", "", "The name of the email contact.")
	notificationscontacts_createEmailContactCmd.Flags().String("tags", "", "A map of tags assigned to a resource.")
	notificationscontacts_createEmailContactCmd.MarkFlagRequired("email-address")
	notificationscontacts_createEmailContactCmd.MarkFlagRequired("name")
	notificationscontactsCmd.AddCommand(notificationscontacts_createEmailContactCmd)
}
