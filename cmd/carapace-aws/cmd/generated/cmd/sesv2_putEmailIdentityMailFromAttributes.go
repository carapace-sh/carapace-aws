package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putEmailIdentityMailFromAttributesCmd = &cobra.Command{
	Use:   "put-email-identity-mail-from-attributes",
	Short: "Used to enable or disable the custom Mail-From domain configuration for an email identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putEmailIdentityMailFromAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_putEmailIdentityMailFromAttributesCmd).Standalone()

		sesv2_putEmailIdentityMailFromAttributesCmd.Flags().String("behavior-on-mx-failure", "", "The action to take if the required MX record isn't found when you send an email.")
		sesv2_putEmailIdentityMailFromAttributesCmd.Flags().String("email-identity", "", "The verified email identity.")
		sesv2_putEmailIdentityMailFromAttributesCmd.Flags().String("mail-from-domain", "", "The custom MAIL FROM domain that you want the verified identity to use.")
		sesv2_putEmailIdentityMailFromAttributesCmd.MarkFlagRequired("email-identity")
	})
	sesv2Cmd.AddCommand(sesv2_putEmailIdentityMailFromAttributesCmd)
}
