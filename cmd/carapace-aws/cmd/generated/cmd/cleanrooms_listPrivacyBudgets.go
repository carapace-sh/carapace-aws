package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listPrivacyBudgetsCmd = &cobra.Command{
	Use:   "list-privacy-budgets",
	Short: "Returns detailed information about the privacy budgets in a specified membership.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listPrivacyBudgetsCmd).Standalone()

	cleanrooms_listPrivacyBudgetsCmd.Flags().String("access-budget-resource-arn", "", "The Amazon Resource Name (ARN) of the access budget resource to filter privacy budgets by.")
	cleanrooms_listPrivacyBudgetsCmd.Flags().String("max-results", "", "The maximum number of results that are returned for an API request call.")
	cleanrooms_listPrivacyBudgetsCmd.Flags().String("membership-identifier", "", "A unique identifier for one of your memberships for a collaboration.")
	cleanrooms_listPrivacyBudgetsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	cleanrooms_listPrivacyBudgetsCmd.Flags().String("privacy-budget-type", "", "The privacy budget type.")
	cleanrooms_listPrivacyBudgetsCmd.MarkFlagRequired("membership-identifier")
	cleanrooms_listPrivacyBudgetsCmd.MarkFlagRequired("privacy-budget-type")
	cleanroomsCmd.AddCommand(cleanrooms_listPrivacyBudgetsCmd)
}
