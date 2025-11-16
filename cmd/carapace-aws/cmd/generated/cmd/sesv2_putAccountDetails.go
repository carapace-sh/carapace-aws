package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putAccountDetailsCmd = &cobra.Command{
	Use:   "put-account-details",
	Short: "Update your Amazon SES account details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putAccountDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_putAccountDetailsCmd).Standalone()

		sesv2_putAccountDetailsCmd.Flags().String("additional-contact-email-addresses", "", "Additional email addresses that you would like to be notified regarding Amazon SES matters.")
		sesv2_putAccountDetailsCmd.Flags().String("contact-language", "", "The language you would prefer to be contacted with.")
		sesv2_putAccountDetailsCmd.Flags().String("mail-type", "", "The type of email your account will send.")
		sesv2_putAccountDetailsCmd.Flags().String("production-access-enabled", "", "Indicates whether or not your account should have production access in the current Amazon Web Services Region.")
		sesv2_putAccountDetailsCmd.Flags().String("use-case-description", "", "A description of the types of email that you plan to send.")
		sesv2_putAccountDetailsCmd.Flags().String("website-url", "", "The URL of your website.")
		sesv2_putAccountDetailsCmd.MarkFlagRequired("mail-type")
		sesv2_putAccountDetailsCmd.MarkFlagRequired("website-url")
	})
	sesv2Cmd.AddCommand(sesv2_putAccountDetailsCmd)
}
