package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_updatePrivacyBudgetTemplateCmd = &cobra.Command{
	Use:   "update-privacy-budget-template",
	Short: "Updates the privacy budget template for the specified collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_updatePrivacyBudgetTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_updatePrivacyBudgetTemplateCmd).Standalone()

		cleanrooms_updatePrivacyBudgetTemplateCmd.Flags().String("membership-identifier", "", "A unique identifier for one of your memberships for a collaboration.")
		cleanrooms_updatePrivacyBudgetTemplateCmd.Flags().String("parameters", "", "Specifies the epsilon and noise parameters for the privacy budget template.")
		cleanrooms_updatePrivacyBudgetTemplateCmd.Flags().String("privacy-budget-template-identifier", "", "A unique identifier for your privacy budget template that you want to update.")
		cleanrooms_updatePrivacyBudgetTemplateCmd.Flags().String("privacy-budget-type", "", "Specifies the type of the privacy budget template.")
		cleanrooms_updatePrivacyBudgetTemplateCmd.MarkFlagRequired("membership-identifier")
		cleanrooms_updatePrivacyBudgetTemplateCmd.MarkFlagRequired("privacy-budget-template-identifier")
		cleanrooms_updatePrivacyBudgetTemplateCmd.MarkFlagRequired("privacy-budget-type")
	})
	cleanroomsCmd.AddCommand(cleanrooms_updatePrivacyBudgetTemplateCmd)
}
