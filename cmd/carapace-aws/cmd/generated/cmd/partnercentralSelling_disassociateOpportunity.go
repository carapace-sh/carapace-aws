package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_disassociateOpportunityCmd = &cobra.Command{
	Use:   "disassociate-opportunity",
	Short: "Allows you to remove an existing association between an `Opportunity` and related entities, such as a Partner Solution, Amazon Web Services product, or an Amazon Web Services Marketplace offer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_disassociateOpportunityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(partnercentralSelling_disassociateOpportunityCmd).Standalone()

		partnercentralSelling_disassociateOpportunityCmd.Flags().String("catalog", "", "Specifies the catalog associated with the request.")
		partnercentralSelling_disassociateOpportunityCmd.Flags().String("opportunity-identifier", "", "The opportunity's unique identifier for when you want to disassociate it from related entities.")
		partnercentralSelling_disassociateOpportunityCmd.Flags().String("related-entity-identifier", "", "The related entity's identifier that you want to disassociate from the opportunity.")
		partnercentralSelling_disassociateOpportunityCmd.Flags().String("related-entity-type", "", "The type of the entity that you're disassociating from the opportunity.")
		partnercentralSelling_disassociateOpportunityCmd.MarkFlagRequired("catalog")
		partnercentralSelling_disassociateOpportunityCmd.MarkFlagRequired("opportunity-identifier")
		partnercentralSelling_disassociateOpportunityCmd.MarkFlagRequired("related-entity-identifier")
		partnercentralSelling_disassociateOpportunityCmd.MarkFlagRequired("related-entity-type")
	})
	partnercentralSellingCmd.AddCommand(partnercentralSelling_disassociateOpportunityCmd)
}
