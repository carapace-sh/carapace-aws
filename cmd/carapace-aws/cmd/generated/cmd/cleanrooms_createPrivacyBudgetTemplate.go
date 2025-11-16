package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_createPrivacyBudgetTemplateCmd = &cobra.Command{
	Use:   "create-privacy-budget-template",
	Short: "Creates a privacy budget template for a specified collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_createPrivacyBudgetTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_createPrivacyBudgetTemplateCmd).Standalone()

		cleanrooms_createPrivacyBudgetTemplateCmd.Flags().String("auto-refresh", "", "How often the privacy budget refreshes.")
		cleanrooms_createPrivacyBudgetTemplateCmd.Flags().String("membership-identifier", "", "A unique identifier for one of your memberships for a collaboration.")
		cleanrooms_createPrivacyBudgetTemplateCmd.Flags().String("parameters", "", "Specifies your parameters for the privacy budget template.")
		cleanrooms_createPrivacyBudgetTemplateCmd.Flags().String("privacy-budget-type", "", "Specifies the type of the privacy budget template.")
		cleanrooms_createPrivacyBudgetTemplateCmd.Flags().String("tags", "", "An optional label that you can assign to a resource when you create it.")
		cleanrooms_createPrivacyBudgetTemplateCmd.MarkFlagRequired("membership-identifier")
		cleanrooms_createPrivacyBudgetTemplateCmd.MarkFlagRequired("parameters")
		cleanrooms_createPrivacyBudgetTemplateCmd.MarkFlagRequired("privacy-budget-type")
	})
	cleanroomsCmd.AddCommand(cleanrooms_createPrivacyBudgetTemplateCmd)
}
