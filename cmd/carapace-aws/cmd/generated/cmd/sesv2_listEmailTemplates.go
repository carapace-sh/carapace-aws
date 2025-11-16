package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_listEmailTemplatesCmd = &cobra.Command{
	Use:   "list-email-templates",
	Short: "Lists the email templates present in your Amazon SES account in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_listEmailTemplatesCmd).Standalone()

	sesv2_listEmailTemplatesCmd.Flags().String("next-token", "", "A token returned from a previous call to `ListEmailTemplates` to indicate the position in the list of email templates.")
	sesv2_listEmailTemplatesCmd.Flags().String("page-size", "", "The number of results to show in a single call to `ListEmailTemplates`.")
	sesv2Cmd.AddCommand(sesv2_listEmailTemplatesCmd)
}
