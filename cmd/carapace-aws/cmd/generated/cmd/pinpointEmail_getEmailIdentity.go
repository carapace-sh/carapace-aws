package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_getEmailIdentityCmd = &cobra.Command{
	Use:   "get-email-identity",
	Short: "Provides information about a specific identity associated with your Amazon Pinpoint account, including the identity's verification status, its DKIM authentication status, and its custom Mail-From settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_getEmailIdentityCmd).Standalone()

	pinpointEmail_getEmailIdentityCmd.Flags().String("email-identity", "", "The email identity that you want to retrieve details for.")
	pinpointEmail_getEmailIdentityCmd.MarkFlagRequired("email-identity")
	pinpointEmailCmd.AddCommand(pinpointEmail_getEmailIdentityCmd)
}
