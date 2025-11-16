package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_putEmailIdentityDkimAttributesCmd = &cobra.Command{
	Use:   "put-email-identity-dkim-attributes",
	Short: "Used to enable or disable DKIM authentication for an email identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_putEmailIdentityDkimAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointEmail_putEmailIdentityDkimAttributesCmd).Standalone()

		pinpointEmail_putEmailIdentityDkimAttributesCmd.Flags().String("email-identity", "", "The email identity that you want to change the DKIM settings for.")
		pinpointEmail_putEmailIdentityDkimAttributesCmd.Flags().String("signing-enabled", "", "Sets the DKIM signing configuration for the identity.")
		pinpointEmail_putEmailIdentityDkimAttributesCmd.MarkFlagRequired("email-identity")
	})
	pinpointEmailCmd.AddCommand(pinpointEmail_putEmailIdentityDkimAttributesCmd)
}
