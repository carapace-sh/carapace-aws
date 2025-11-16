package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getPrivacyBudgetTemplateCmd = &cobra.Command{
	Use:   "get-privacy-budget-template",
	Short: "Returns details for a specified privacy budget template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getPrivacyBudgetTemplateCmd).Standalone()

	cleanrooms_getPrivacyBudgetTemplateCmd.Flags().String("membership-identifier", "", "A unique identifier for one of your memberships for a collaboration.")
	cleanrooms_getPrivacyBudgetTemplateCmd.Flags().String("privacy-budget-template-identifier", "", "A unique identifier for your privacy budget template.")
	cleanrooms_getPrivacyBudgetTemplateCmd.MarkFlagRequired("membership-identifier")
	cleanrooms_getPrivacyBudgetTemplateCmd.MarkFlagRequired("privacy-budget-template-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_getPrivacyBudgetTemplateCmd)
}
