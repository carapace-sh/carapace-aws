package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteAccountCustomizationCmd = &cobra.Command{
	Use:   "delete-account-customization",
	Short: "This API permanently deletes all Quick Sight customizations for the specified Amazon Web Services account and namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteAccountCustomizationCmd).Standalone()

	quicksight_deleteAccountCustomizationCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that you want to delete Quick Sight customizations from.")
	quicksight_deleteAccountCustomizationCmd.Flags().String("namespace", "", "The Quick Sight namespace that you're deleting the customizations from.")
	quicksight_deleteAccountCustomizationCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_deleteAccountCustomizationCmd)
}
