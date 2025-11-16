package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_listEngagementMembersCmd = &cobra.Command{
	Use:   "list-engagement-members",
	Short: "Retrieves the details of member partners in an Engagement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_listEngagementMembersCmd).Standalone()

	partnercentralSelling_listEngagementMembersCmd.Flags().String("catalog", "", "The catalog related to the request.")
	partnercentralSelling_listEngagementMembersCmd.Flags().String("identifier", "", "Identifier of the Engagement record to retrieve members from.")
	partnercentralSelling_listEngagementMembersCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	partnercentralSelling_listEngagementMembersCmd.Flags().String("next-token", "", "The token for the next set of results.")
	partnercentralSelling_listEngagementMembersCmd.MarkFlagRequired("catalog")
	partnercentralSelling_listEngagementMembersCmd.MarkFlagRequired("identifier")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_listEngagementMembersCmd)
}
