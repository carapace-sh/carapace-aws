package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateCustomPermissionsCmd = &cobra.Command{
	Use:   "update-custom-permissions",
	Short: "Updates a custom permissions profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateCustomPermissionsCmd).Standalone()

	quicksight_updateCustomPermissionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the custom permissions profile that you want to update.")
	quicksight_updateCustomPermissionsCmd.Flags().String("capabilities", "", "A set of actions to include in the custom permissions profile.")
	quicksight_updateCustomPermissionsCmd.Flags().String("custom-permissions-name", "", "The name of the custom permissions profile that you want to update.")
	quicksight_updateCustomPermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateCustomPermissionsCmd.MarkFlagRequired("custom-permissions-name")
	quicksightCmd.AddCommand(quicksight_updateCustomPermissionsCmd)
}
