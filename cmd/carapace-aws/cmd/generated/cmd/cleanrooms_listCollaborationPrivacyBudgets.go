package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listCollaborationPrivacyBudgetsCmd = &cobra.Command{
	Use:   "list-collaboration-privacy-budgets",
	Short: "Returns an array that summarizes each privacy budget in a specified collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listCollaborationPrivacyBudgetsCmd).Standalone()

	cleanrooms_listCollaborationPrivacyBudgetsCmd.Flags().String("access-budget-resource-arn", "", "The Amazon Resource Name (ARN) of the Configured Table Association (ConfiguredTableAssociation) used to filter privacy budgets.")
	cleanrooms_listCollaborationPrivacyBudgetsCmd.Flags().String("collaboration-identifier", "", "A unique identifier for one of your collaborations.")
	cleanrooms_listCollaborationPrivacyBudgetsCmd.Flags().String("max-results", "", "The maximum number of results that are returned for an API request call.")
	cleanrooms_listCollaborationPrivacyBudgetsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	cleanrooms_listCollaborationPrivacyBudgetsCmd.Flags().String("privacy-budget-type", "", "Specifies the type of the privacy budget.")
	cleanrooms_listCollaborationPrivacyBudgetsCmd.MarkFlagRequired("collaboration-identifier")
	cleanrooms_listCollaborationPrivacyBudgetsCmd.MarkFlagRequired("privacy-budget-type")
	cleanroomsCmd.AddCommand(cleanrooms_listCollaborationPrivacyBudgetsCmd)
}
