package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_createEmailIdentityCmd = &cobra.Command{
	Use:   "create-email-identity",
	Short: "Starts the process of verifying an email identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_createEmailIdentityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_createEmailIdentityCmd).Standalone()

		sesv2_createEmailIdentityCmd.Flags().String("configuration-set-name", "", "The configuration set to use by default when sending from this identity.")
		sesv2_createEmailIdentityCmd.Flags().String("dkim-signing-attributes", "", "If your request includes this object, Amazon SES configures the identity to use Bring Your Own DKIM (BYODKIM) for DKIM authentication purposes, or, configures the key length to be used for [Easy DKIM](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html).")
		sesv2_createEmailIdentityCmd.Flags().String("email-identity", "", "The email address or domain to verify.")
		sesv2_createEmailIdentityCmd.Flags().String("tags", "", "An array of objects that define the tags (keys and values) to associate with the email identity.")
		sesv2_createEmailIdentityCmd.MarkFlagRequired("email-identity")
	})
	sesv2Cmd.AddCommand(sesv2_createEmailIdentityCmd)
}
