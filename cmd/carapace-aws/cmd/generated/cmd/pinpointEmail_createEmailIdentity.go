package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_createEmailIdentityCmd = &cobra.Command{
	Use:   "create-email-identity",
	Short: "Verifies an email identity for use with Amazon Pinpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_createEmailIdentityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointEmail_createEmailIdentityCmd).Standalone()

		pinpointEmail_createEmailIdentityCmd.Flags().String("email-identity", "", "The email address or domain that you want to verify.")
		pinpointEmail_createEmailIdentityCmd.Flags().String("tags", "", "An array of objects that define the tags (keys and values) that you want to associate with the email identity.")
		pinpointEmail_createEmailIdentityCmd.MarkFlagRequired("email-identity")
	})
	pinpointEmailCmd.AddCommand(pinpointEmail_createEmailIdentityCmd)
}
