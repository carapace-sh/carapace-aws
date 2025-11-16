package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateAccountCustomizationCmd = &cobra.Command{
	Use:   "update-account-customization",
	Short: "Updates Amazon Quick Sight customizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateAccountCustomizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateAccountCustomizationCmd).Standalone()

		quicksight_updateAccountCustomizationCmd.Flags().String("account-customization", "", "The Quick Sight customizations you're updating.")
		quicksight_updateAccountCustomizationCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that you want to update Quick Sight customizations for.")
		quicksight_updateAccountCustomizationCmd.Flags().String("namespace", "", "The namespace that you want to update Quick Sight customizations for.")
		quicksight_updateAccountCustomizationCmd.MarkFlagRequired("account-customization")
		quicksight_updateAccountCustomizationCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_updateAccountCustomizationCmd)
}
