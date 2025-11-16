package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_startPrimaryEmailUpdateCmd = &cobra.Command{
	Use:   "start-primary-email-update",
	Short: "Starts the process to update the primary email address for the specified account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_startPrimaryEmailUpdateCmd).Standalone()

	account_startPrimaryEmailUpdateCmd.Flags().String("account-id", "", "Specifies the 12-digit account ID number of the Amazon Web Services account that you want to access or modify with this operation.")
	account_startPrimaryEmailUpdateCmd.Flags().String("primary-email", "", "The new primary email address (also known as the root user email address) to use in the specified account.")
	account_startPrimaryEmailUpdateCmd.MarkFlagRequired("account-id")
	account_startPrimaryEmailUpdateCmd.MarkFlagRequired("primary-email")
	accountCmd.AddCommand(account_startPrimaryEmailUpdateCmd)
}
