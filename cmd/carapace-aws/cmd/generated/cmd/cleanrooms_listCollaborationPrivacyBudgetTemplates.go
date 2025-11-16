package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listCollaborationPrivacyBudgetTemplatesCmd = &cobra.Command{
	Use:   "list-collaboration-privacy-budget-templates",
	Short: "Returns an array that summarizes each privacy budget template in a specified collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listCollaborationPrivacyBudgetTemplatesCmd).Standalone()

	cleanrooms_listCollaborationPrivacyBudgetTemplatesCmd.Flags().String("collaboration-identifier", "", "A unique identifier for one of your collaborations.")
	cleanrooms_listCollaborationPrivacyBudgetTemplatesCmd.Flags().String("max-results", "", "The maximum number of results that are returned for an API request call.")
	cleanrooms_listCollaborationPrivacyBudgetTemplatesCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	cleanrooms_listCollaborationPrivacyBudgetTemplatesCmd.MarkFlagRequired("collaboration-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_listCollaborationPrivacyBudgetTemplatesCmd)
}
