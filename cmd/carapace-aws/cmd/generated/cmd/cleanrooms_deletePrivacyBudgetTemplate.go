package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_deletePrivacyBudgetTemplateCmd = &cobra.Command{
	Use:   "delete-privacy-budget-template",
	Short: "Deletes a privacy budget template for a specified collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_deletePrivacyBudgetTemplateCmd).Standalone()

	cleanrooms_deletePrivacyBudgetTemplateCmd.Flags().String("membership-identifier", "", "A unique identifier for one of your memberships for a collaboration.")
	cleanrooms_deletePrivacyBudgetTemplateCmd.Flags().String("privacy-budget-template-identifier", "", "A unique identifier for your privacy budget template.")
	cleanrooms_deletePrivacyBudgetTemplateCmd.MarkFlagRequired("membership-identifier")
	cleanrooms_deletePrivacyBudgetTemplateCmd.MarkFlagRequired("privacy-budget-template-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_deletePrivacyBudgetTemplateCmd)
}
