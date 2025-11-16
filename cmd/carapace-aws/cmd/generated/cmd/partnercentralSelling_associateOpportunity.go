package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_associateOpportunityCmd = &cobra.Command{
	Use:   "associate-opportunity",
	Short: "Enables you to create a formal association between an `Opportunity` and various related entities, enriching the context and details of the opportunity for better collaboration and decision making.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_associateOpportunityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(partnercentralSelling_associateOpportunityCmd).Standalone()

		partnercentralSelling_associateOpportunityCmd.Flags().String("catalog", "", "Specifies the catalog associated with the request.")
		partnercentralSelling_associateOpportunityCmd.Flags().String("opportunity-identifier", "", "Requires the `Opportunity`'s unique identifier when you want to associate it with a related entity.")
		partnercentralSelling_associateOpportunityCmd.Flags().String("related-entity-identifier", "", "Requires the related entity's unique identifier when you want to associate it with the `Opportunity`.")
		partnercentralSelling_associateOpportunityCmd.Flags().String("related-entity-type", "", "Specifies the entity type that you're associating with the `Opportunity`.")
		partnercentralSelling_associateOpportunityCmd.MarkFlagRequired("catalog")
		partnercentralSelling_associateOpportunityCmd.MarkFlagRequired("opportunity-identifier")
		partnercentralSelling_associateOpportunityCmd.MarkFlagRequired("related-entity-identifier")
		partnercentralSelling_associateOpportunityCmd.MarkFlagRequired("related-entity-type")
	})
	partnercentralSellingCmd.AddCommand(partnercentralSelling_associateOpportunityCmd)
}
