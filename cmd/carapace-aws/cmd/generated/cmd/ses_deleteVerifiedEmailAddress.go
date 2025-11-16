package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_deleteVerifiedEmailAddressCmd = &cobra.Command{
	Use:   "delete-verified-email-address",
	Short: "Deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_deleteVerifiedEmailAddressCmd).Standalone()

	ses_deleteVerifiedEmailAddressCmd.Flags().String("email-address", "", "An email address to be removed from the list of verified addresses.")
	ses_deleteVerifiedEmailAddressCmd.MarkFlagRequired("email-address")
	sesCmd.AddCommand(ses_deleteVerifiedEmailAddressCmd)
}
