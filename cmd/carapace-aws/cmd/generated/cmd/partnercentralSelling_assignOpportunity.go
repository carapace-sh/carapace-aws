package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_assignOpportunityCmd = &cobra.Command{
	Use:   "assign-opportunity",
	Short: "Enables you to reassign an existing `Opportunity` to another user within your Partner Central account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_assignOpportunityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(partnercentralSelling_assignOpportunityCmd).Standalone()

		partnercentralSelling_assignOpportunityCmd.Flags().String("assignee", "", "Specifies the user or team member responsible for managing the assigned opportunity.")
		partnercentralSelling_assignOpportunityCmd.Flags().String("catalog", "", "Specifies the catalog associated with the request.")
		partnercentralSelling_assignOpportunityCmd.Flags().String("identifier", "", "Requires the `Opportunity`'s unique identifier when you want to assign it to another user.")
		partnercentralSelling_assignOpportunityCmd.MarkFlagRequired("assignee")
		partnercentralSelling_assignOpportunityCmd.MarkFlagRequired("catalog")
		partnercentralSelling_assignOpportunityCmd.MarkFlagRequired("identifier")
	})
	partnercentralSellingCmd.AddCommand(partnercentralSelling_assignOpportunityCmd)
}
