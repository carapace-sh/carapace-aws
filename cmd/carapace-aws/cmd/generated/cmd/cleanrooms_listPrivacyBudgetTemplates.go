package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listPrivacyBudgetTemplatesCmd = &cobra.Command{
	Use:   "list-privacy-budget-templates",
	Short: "Returns detailed information about the privacy budget templates in a specified membership.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listPrivacyBudgetTemplatesCmd).Standalone()

	cleanrooms_listPrivacyBudgetTemplatesCmd.Flags().String("max-results", "", "The maximum number of results that are returned for an API request call.")
	cleanrooms_listPrivacyBudgetTemplatesCmd.Flags().String("membership-identifier", "", "A unique identifier for one of your memberships for a collaboration.")
	cleanrooms_listPrivacyBudgetTemplatesCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	cleanrooms_listPrivacyBudgetTemplatesCmd.MarkFlagRequired("membership-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_listPrivacyBudgetTemplatesCmd)
}
