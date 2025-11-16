package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createAccountCustomizationCmd = &cobra.Command{
	Use:   "create-account-customization",
	Short: "Creates Amazon Quick Sight customizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createAccountCustomizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_createAccountCustomizationCmd).Standalone()

		quicksight_createAccountCustomizationCmd.Flags().String("account-customization", "", "The Quick Sight customizations you're adding.")
		quicksight_createAccountCustomizationCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that you want to customize Quick Sight for.")
		quicksight_createAccountCustomizationCmd.Flags().String("namespace", "", "The Quick Sight namespace that you want to add customizations to.")
		quicksight_createAccountCustomizationCmd.Flags().String("tags", "", "A list of the tags that you want to attach to this resource.")
		quicksight_createAccountCustomizationCmd.MarkFlagRequired("account-customization")
		quicksight_createAccountCustomizationCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_createAccountCustomizationCmd)
}
