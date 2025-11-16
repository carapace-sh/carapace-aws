package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_updateOpportunityCmd = &cobra.Command{
	Use:   "update-opportunity",
	Short: "Updates the `Opportunity` record identified by a given `Identifier`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_updateOpportunityCmd).Standalone()

	partnercentralSelling_updateOpportunityCmd.Flags().String("catalog", "", "Specifies the catalog associated with the request.")
	partnercentralSelling_updateOpportunityCmd.Flags().String("customer", "", "Specifies details of the customer associated with the `Opportunity`.")
	partnercentralSelling_updateOpportunityCmd.Flags().String("identifier", "", "Read-only, system generated `Opportunity` unique identifier.")
	partnercentralSelling_updateOpportunityCmd.Flags().String("last-modified-date", "", "`DateTime` when the opportunity was last modified.")
	partnercentralSelling_updateOpportunityCmd.Flags().String("life-cycle", "", "An object that contains lifecycle details for the `Opportunity`.")
	partnercentralSelling_updateOpportunityCmd.Flags().String("marketing", "", "An object that contains marketing details for the `Opportunity`.")
	partnercentralSelling_updateOpportunityCmd.Flags().String("national-security", "", "Specifies if the opportunity is associated with national security concerns.")
	partnercentralSelling_updateOpportunityCmd.Flags().String("opportunity-type", "", "Specifies the opportunity type as a renewal, new, or expansion.")
	partnercentralSelling_updateOpportunityCmd.Flags().String("partner-opportunity-identifier", "", "Specifies the opportunity's unique identifier in the partner's CRM system.")
	partnercentralSelling_updateOpportunityCmd.Flags().String("primary-needs-from-aws", "", "Identifies the type of support the partner needs from Amazon Web Services.")
	partnercentralSelling_updateOpportunityCmd.Flags().String("project", "", "An object that contains project details summary for the `Opportunity`.")
	partnercentralSelling_updateOpportunityCmd.Flags().String("software-revenue", "", "Specifies details of a customer's procurement terms.")
	partnercentralSelling_updateOpportunityCmd.MarkFlagRequired("catalog")
	partnercentralSelling_updateOpportunityCmd.MarkFlagRequired("identifier")
	partnercentralSelling_updateOpportunityCmd.MarkFlagRequired("last-modified-date")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_updateOpportunityCmd)
}
