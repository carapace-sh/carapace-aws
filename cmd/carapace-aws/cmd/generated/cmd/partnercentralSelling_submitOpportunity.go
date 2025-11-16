package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_submitOpportunityCmd = &cobra.Command{
	Use:   "submit-opportunity",
	Short: "Use this action to submit an Opportunity that was previously created by partner for AWS review.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_submitOpportunityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(partnercentralSelling_submitOpportunityCmd).Standalone()

		partnercentralSelling_submitOpportunityCmd.Flags().String("catalog", "", "Specifies the catalog related to the request.")
		partnercentralSelling_submitOpportunityCmd.Flags().String("identifier", "", "The identifier of the Opportunity previously created by partner and needs to be submitted.")
		partnercentralSelling_submitOpportunityCmd.Flags().String("involvement-type", "", "Specifies the level of AWS sellers' involvement on the opportunity.")
		partnercentralSelling_submitOpportunityCmd.Flags().String("visibility", "", "Determines whether to restrict visibility of the opportunity from AWS sales.")
		partnercentralSelling_submitOpportunityCmd.MarkFlagRequired("catalog")
		partnercentralSelling_submitOpportunityCmd.MarkFlagRequired("identifier")
		partnercentralSelling_submitOpportunityCmd.MarkFlagRequired("involvement-type")
	})
	partnercentralSellingCmd.AddCommand(partnercentralSelling_submitOpportunityCmd)
}
