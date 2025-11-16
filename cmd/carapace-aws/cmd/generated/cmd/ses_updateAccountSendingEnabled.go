package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_updateAccountSendingEnabledCmd = &cobra.Command{
	Use:   "update-account-sending-enabled",
	Short: "Enables or disables email sending across your entire Amazon SES account in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_updateAccountSendingEnabledCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_updateAccountSendingEnabledCmd).Standalone()

		ses_updateAccountSendingEnabledCmd.Flags().String("enabled", "", "Describes whether email sending is enabled or disabled for your Amazon SES account in the current Amazon Web Services Region.")
	})
	sesCmd.AddCommand(ses_updateAccountSendingEnabledCmd)
}
