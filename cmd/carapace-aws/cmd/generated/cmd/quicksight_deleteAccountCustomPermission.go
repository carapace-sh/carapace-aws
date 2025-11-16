package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteAccountCustomPermissionCmd = &cobra.Command{
	Use:   "delete-account-custom-permission",
	Short: "Unapplies a custom permissions profile from an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteAccountCustomPermissionCmd).Standalone()

	quicksight_deleteAccountCustomPermissionCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account from which you want to unapply the custom permissions profile.")
	quicksight_deleteAccountCustomPermissionCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_deleteAccountCustomPermissionCmd)
}
