package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_listSolutionsCmd = &cobra.Command{
	Use:   "list-solutions",
	Short: "Retrieves a list of Partner Solutions that the partner registered on Partner Central.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_listSolutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(partnercentralSelling_listSolutionsCmd).Standalone()

		partnercentralSelling_listSolutionsCmd.Flags().String("catalog", "", "Specifies the catalog associated with the request.")
		partnercentralSelling_listSolutionsCmd.Flags().String("category", "", "Filters the solutions based on the category to which they belong.")
		partnercentralSelling_listSolutionsCmd.Flags().String("identifier", "", "Filters the solutions based on their unique identifier.")
		partnercentralSelling_listSolutionsCmd.Flags().String("max-results", "", "The maximum number of results returned by a single call.")
		partnercentralSelling_listSolutionsCmd.Flags().String("next-token", "", "A pagination token used to retrieve the next set of results in subsequent calls.")
		partnercentralSelling_listSolutionsCmd.Flags().String("sort", "", "Object that configures sorting done on the response.")
		partnercentralSelling_listSolutionsCmd.Flags().String("status", "", "Filters solutions based on their status.")
		partnercentralSelling_listSolutionsCmd.MarkFlagRequired("catalog")
	})
	partnercentralSellingCmd.AddCommand(partnercentralSelling_listSolutionsCmd)
}
