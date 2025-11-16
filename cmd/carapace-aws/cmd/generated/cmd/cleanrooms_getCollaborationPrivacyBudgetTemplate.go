package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getCollaborationPrivacyBudgetTemplateCmd = &cobra.Command{
	Use:   "get-collaboration-privacy-budget-template",
	Short: "Returns details about a specified privacy budget template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getCollaborationPrivacyBudgetTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_getCollaborationPrivacyBudgetTemplateCmd).Standalone()

		cleanrooms_getCollaborationPrivacyBudgetTemplateCmd.Flags().String("collaboration-identifier", "", "A unique identifier for one of your collaborations.")
		cleanrooms_getCollaborationPrivacyBudgetTemplateCmd.Flags().String("privacy-budget-template-identifier", "", "A unique identifier for one of your privacy budget templates.")
		cleanrooms_getCollaborationPrivacyBudgetTemplateCmd.MarkFlagRequired("collaboration-identifier")
		cleanrooms_getCollaborationPrivacyBudgetTemplateCmd.MarkFlagRequired("privacy-budget-template-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_getCollaborationPrivacyBudgetTemplateCmd)
}
