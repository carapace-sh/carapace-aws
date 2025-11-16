package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_getAccountSendingEnabledCmd = &cobra.Command{
	Use:   "get-account-sending-enabled",
	Short: "Returns the email sending status of the Amazon SES account for the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_getAccountSendingEnabledCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_getAccountSendingEnabledCmd).Standalone()

	})
	sesCmd.AddCommand(ses_getAccountSendingEnabledCmd)
}
