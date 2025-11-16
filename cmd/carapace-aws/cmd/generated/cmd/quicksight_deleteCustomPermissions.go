package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteCustomPermissionsCmd = &cobra.Command{
	Use:   "delete-custom-permissions",
	Short: "Deletes a custom permissions profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteCustomPermissionsCmd).Standalone()

	quicksight_deleteCustomPermissionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the custom permissions profile that you want to delete.")
	quicksight_deleteCustomPermissionsCmd.Flags().String("custom-permissions-name", "", "The name of the custom permissions profile that you want to delete.")
	quicksight_deleteCustomPermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksight_deleteCustomPermissionsCmd.MarkFlagRequired("custom-permissions-name")
	quicksightCmd.AddCommand(quicksight_deleteCustomPermissionsCmd)
}
