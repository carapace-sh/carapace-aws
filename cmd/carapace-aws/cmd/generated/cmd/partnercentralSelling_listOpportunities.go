package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_listOpportunitiesCmd = &cobra.Command{
	Use:   "list-opportunities",
	Short: "This request accepts a list of filters that retrieve opportunity subsets as well as sort options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_listOpportunitiesCmd).Standalone()

	partnercentralSelling_listOpportunitiesCmd.Flags().String("catalog", "", "Specifies the catalog associated with the request.")
	partnercentralSelling_listOpportunitiesCmd.Flags().String("customer-company-name", "", "Filters the opportunities based on the customer's company name.")
	partnercentralSelling_listOpportunitiesCmd.Flags().String("identifier", "", "Filters the opportunities based on the opportunity identifier.")
	partnercentralSelling_listOpportunitiesCmd.Flags().String("last-modified-date", "", "Filters the opportunities based on their last modified date.")
	partnercentralSelling_listOpportunitiesCmd.Flags().String("life-cycle-review-status", "", "Filters the opportunities based on their current lifecycle approval status.")
	partnercentralSelling_listOpportunitiesCmd.Flags().String("life-cycle-stage", "", "Filters the opportunities based on their lifecycle stage.")
	partnercentralSelling_listOpportunitiesCmd.Flags().String("max-results", "", "Specifies the maximum number of results to return in a single call.")
	partnercentralSelling_listOpportunitiesCmd.Flags().String("next-token", "", "A pagination token used to retrieve the next set of results in subsequent calls.")
	partnercentralSelling_listOpportunitiesCmd.Flags().String("sort", "", "An object that specifies how the response is sorted.")
	partnercentralSelling_listOpportunitiesCmd.MarkFlagRequired("catalog")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_listOpportunitiesCmd)
}
