package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_listTemplatesCmd = &cobra.Command{
	Use:   "list-templates",
	Short: "Lists the email templates present in your Amazon SES account in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_listTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_listTemplatesCmd).Standalone()

		ses_listTemplatesCmd.Flags().String("max-items", "", "The maximum number of templates to return.")
		ses_listTemplatesCmd.Flags().String("next-token", "", "A token returned from a previous call to `ListTemplates` to indicate the position in the list of email templates.")
	})
	sesCmd.AddCommand(ses_listTemplatesCmd)
}
