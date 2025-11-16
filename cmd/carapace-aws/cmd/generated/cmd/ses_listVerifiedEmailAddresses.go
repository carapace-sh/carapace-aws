package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_listVerifiedEmailAddressesCmd = &cobra.Command{
	Use:   "list-verified-email-addresses",
	Short: "Deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_listVerifiedEmailAddressesCmd).Standalone()

	sesCmd.AddCommand(ses_listVerifiedEmailAddressesCmd)
}
