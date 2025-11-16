package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateAccountCustomPermissionCmd = &cobra.Command{
	Use:   "update-account-custom-permission",
	Short: "Applies a custom permissions profile to an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateAccountCustomPermissionCmd).Standalone()

	quicksight_updateAccountCustomPermissionCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account for which you want to apply a custom permissions profile.")
	quicksight_updateAccountCustomPermissionCmd.Flags().String("custom-permissions-name", "", "The name of the custom permissions profile that you want to apply to an account.")
	quicksight_updateAccountCustomPermissionCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateAccountCustomPermissionCmd.MarkFlagRequired("custom-permissions-name")
	quicksightCmd.AddCommand(quicksight_updateAccountCustomPermissionCmd)
}
