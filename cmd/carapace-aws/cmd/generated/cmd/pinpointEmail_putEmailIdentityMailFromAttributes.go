package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_putEmailIdentityMailFromAttributesCmd = &cobra.Command{
	Use:   "put-email-identity-mail-from-attributes",
	Short: "Used to enable or disable the custom Mail-From domain configuration for an email identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_putEmailIdentityMailFromAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointEmail_putEmailIdentityMailFromAttributesCmd).Standalone()

		pinpointEmail_putEmailIdentityMailFromAttributesCmd.Flags().String("behavior-on-mx-failure", "", "The action that you want Amazon Pinpoint to take if it can't read the required MX record when you send an email.")
		pinpointEmail_putEmailIdentityMailFromAttributesCmd.Flags().String("email-identity", "", "The verified email identity that you want to set up the custom MAIL FROM domain for.")
		pinpointEmail_putEmailIdentityMailFromAttributesCmd.Flags().String("mail-from-domain", "", "The custom MAIL FROM domain that you want the verified identity to use.")
		pinpointEmail_putEmailIdentityMailFromAttributesCmd.MarkFlagRequired("email-identity")
	})
	pinpointEmailCmd.AddCommand(pinpointEmail_putEmailIdentityMailFromAttributesCmd)
}
