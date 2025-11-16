package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_listCallAnalyticsCategoriesCmd = &cobra.Command{
	Use:   "list-call-analytics-categories",
	Short: "Provides a list of Call Analytics categories, including all rules that make up each category.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_listCallAnalyticsCategoriesCmd).Standalone()

	transcribe_listCallAnalyticsCategoriesCmd.Flags().String("max-results", "", "The maximum number of Call Analytics categories to return in each page of results.")
	transcribe_listCallAnalyticsCategoriesCmd.Flags().String("next-token", "", "If your `ListCallAnalyticsCategories` request returns more results than can be displayed, `NextToken` is displayed in the response with an associated string.")
	transcribeCmd.AddCommand(transcribe_listCallAnalyticsCategoriesCmd)
}
