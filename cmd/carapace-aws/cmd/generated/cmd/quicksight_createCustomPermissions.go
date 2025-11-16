package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createCustomPermissionsCmd = &cobra.Command{
	Use:   "create-custom-permissions",
	Short: "Creates a custom permissions profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createCustomPermissionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_createCustomPermissionsCmd).Standalone()

		quicksight_createCustomPermissionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that you want to create the custom permissions profile in.")
		quicksight_createCustomPermissionsCmd.Flags().String("capabilities", "", "A set of actions to include in the custom permissions profile.")
		quicksight_createCustomPermissionsCmd.Flags().String("custom-permissions-name", "", "The name of the custom permissions profile that you want to create.")
		quicksight_createCustomPermissionsCmd.Flags().String("tags", "", "The tags to associate with the custom permissions profile.")
		quicksight_createCustomPermissionsCmd.MarkFlagRequired("aws-account-id")
		quicksight_createCustomPermissionsCmd.MarkFlagRequired("custom-permissions-name")
	})
	quicksightCmd.AddCommand(quicksight_createCustomPermissionsCmd)
}
