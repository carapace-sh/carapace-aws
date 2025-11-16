package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_listCustomVerificationEmailTemplatesCmd = &cobra.Command{
	Use:   "list-custom-verification-email-templates",
	Short: "Lists the existing custom verification email templates for your account in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_listCustomVerificationEmailTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_listCustomVerificationEmailTemplatesCmd).Standalone()

		sesv2_listCustomVerificationEmailTemplatesCmd.Flags().String("next-token", "", "A token returned from a previous call to `ListCustomVerificationEmailTemplates` to indicate the position in the list of custom verification email templates.")
		sesv2_listCustomVerificationEmailTemplatesCmd.Flags().String("page-size", "", "The number of results to show in a single call to `ListCustomVerificationEmailTemplates`.")
	})
	sesv2Cmd.AddCommand(sesv2_listCustomVerificationEmailTemplatesCmd)
}
