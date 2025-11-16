package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_getOpportunityCmd = &cobra.Command{
	Use:   "get-opportunity",
	Short: "Fetches the `Opportunity` record from Partner Central by a given `Identifier`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_getOpportunityCmd).Standalone()

	partnercentralSelling_getOpportunityCmd.Flags().String("catalog", "", "Specifies the catalog associated with the request.")
	partnercentralSelling_getOpportunityCmd.Flags().String("identifier", "", "Read-only, system generated `Opportunity` unique identifier.")
	partnercentralSelling_getOpportunityCmd.MarkFlagRequired("catalog")
	partnercentralSelling_getOpportunityCmd.MarkFlagRequired("identifier")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_getOpportunityCmd)
}
