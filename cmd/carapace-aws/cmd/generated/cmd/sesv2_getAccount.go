package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getAccountCmd = &cobra.Command{
	Use:   "get-account",
	Short: "Obtain information about the email-sending status and capabilities of your Amazon SES account in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_getAccountCmd).Standalone()

	})
	sesv2Cmd.AddCommand(sesv2_getAccountCmd)
}
