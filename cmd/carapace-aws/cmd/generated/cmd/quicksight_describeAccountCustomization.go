package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeAccountCustomizationCmd = &cobra.Command{
	Use:   "describe-account-customization",
	Short: "Describes the customizations associated with the provided Amazon Web Services account and Amazon Quick Sight namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeAccountCustomizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_describeAccountCustomizationCmd).Standalone()

		quicksight_describeAccountCustomizationCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that you want to describe Quick Sight customizations for.")
		quicksight_describeAccountCustomizationCmd.Flags().String("namespace", "", "The Quick Sight namespace that you want to describe Quick Sight customizations for.")
		quicksight_describeAccountCustomizationCmd.Flags().String("resolved", "", "The `Resolved` flag works with the other parameters to determine which view of Quick Sight customizations is returned.")
		quicksight_describeAccountCustomizationCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_describeAccountCustomizationCmd)
}
