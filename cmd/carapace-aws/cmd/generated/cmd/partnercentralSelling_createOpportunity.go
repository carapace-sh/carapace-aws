package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_createOpportunityCmd = &cobra.Command{
	Use:   "create-opportunity",
	Short: "Creates an `Opportunity` record in Partner Central.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_createOpportunityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(partnercentralSelling_createOpportunityCmd).Standalone()

		partnercentralSelling_createOpportunityCmd.Flags().String("catalog", "", "Specifies the catalog associated with the request.")
		partnercentralSelling_createOpportunityCmd.Flags().String("client-token", "", "Required to be unique, and should be unchanging, it can be randomly generated or a meaningful string.")
		partnercentralSelling_createOpportunityCmd.Flags().String("customer", "", "Specifies customer details associated with the `Opportunity`.")
		partnercentralSelling_createOpportunityCmd.Flags().String("life-cycle", "", "An object that contains lifecycle details for the `Opportunity`.")
		partnercentralSelling_createOpportunityCmd.Flags().String("marketing", "", "This object contains marketing details and is optional for an opportunity.")
		partnercentralSelling_createOpportunityCmd.Flags().String("national-security", "", "Indicates whether the `Opportunity` pertains to a national security project.")
		partnercentralSelling_createOpportunityCmd.Flags().String("opportunity-team", "", "Represents the internal team handling the opportunity.")
		partnercentralSelling_createOpportunityCmd.Flags().String("opportunity-type", "", "Specifies the opportunity type as a renewal, new, or expansion.")
		partnercentralSelling_createOpportunityCmd.Flags().String("origin", "", "Specifies the origin of the opportunity, indicating if it was sourced from Amazon Web Services or the partner.")
		partnercentralSelling_createOpportunityCmd.Flags().String("partner-opportunity-identifier", "", "Specifies the opportunity's unique identifier in the partner's CRM system.")
		partnercentralSelling_createOpportunityCmd.Flags().String("primary-needs-from-aws", "", "Identifies the type of support the partner needs from Amazon Web Services.")
		partnercentralSelling_createOpportunityCmd.Flags().String("project", "", "An object that contains project details for the `Opportunity`.")
		partnercentralSelling_createOpportunityCmd.Flags().String("software-revenue", "", "Specifies details of a customer's procurement terms.")
		partnercentralSelling_createOpportunityCmd.Flags().String("tags", "", "A map of the key-value pairs of the tag or tags to assign.")
		partnercentralSelling_createOpportunityCmd.MarkFlagRequired("catalog")
		partnercentralSelling_createOpportunityCmd.MarkFlagRequired("client-token")
	})
	partnercentralSellingCmd.AddCommand(partnercentralSelling_createOpportunityCmd)
}
