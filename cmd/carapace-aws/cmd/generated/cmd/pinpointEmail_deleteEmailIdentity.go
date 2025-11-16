package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_deleteEmailIdentityCmd = &cobra.Command{
	Use:   "delete-email-identity",
	Short: "Deletes an email identity that you previously verified for use with Amazon Pinpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_deleteEmailIdentityCmd).Standalone()

	pinpointEmail_deleteEmailIdentityCmd.Flags().String("email-identity", "", "The identity (that is, the email address or domain) that you want to delete from your Amazon Pinpoint account.")
	pinpointEmail_deleteEmailIdentityCmd.MarkFlagRequired("email-identity")
	pinpointEmailCmd.AddCommand(pinpointEmail_deleteEmailIdentityCmd)
}
