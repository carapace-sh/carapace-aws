package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putEmailIdentityDkimSigningAttributesCmd = &cobra.Command{
	Use:   "put-email-identity-dkim-signing-attributes",
	Short: "Used to configure or change the DKIM authentication settings for an email domain identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putEmailIdentityDkimSigningAttributesCmd).Standalone()

	sesv2_putEmailIdentityDkimSigningAttributesCmd.Flags().String("email-identity", "", "The email identity.")
	sesv2_putEmailIdentityDkimSigningAttributesCmd.Flags().String("signing-attributes", "", "An object that contains information about the private key and selector that you want to use to configure DKIM for the identity for Bring Your Own DKIM (BYODKIM) for the identity, or, configures the key length to be used for [Easy DKIM](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html).")
	sesv2_putEmailIdentityDkimSigningAttributesCmd.Flags().String("signing-attributes-origin", "", "The method to use to configure DKIM for the identity.")
	sesv2_putEmailIdentityDkimSigningAttributesCmd.MarkFlagRequired("email-identity")
	sesv2_putEmailIdentityDkimSigningAttributesCmd.MarkFlagRequired("signing-attributes-origin")
	sesv2Cmd.AddCommand(sesv2_putEmailIdentityDkimSigningAttributesCmd)
}
