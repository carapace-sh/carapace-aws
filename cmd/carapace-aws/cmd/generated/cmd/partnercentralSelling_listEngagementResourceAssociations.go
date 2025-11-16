package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_listEngagementResourceAssociationsCmd = &cobra.Command{
	Use:   "list-engagement-resource-associations",
	Short: "Lists the associations between resources and engagements where the caller is a member and has at least one snapshot in the engagement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_listEngagementResourceAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(partnercentralSelling_listEngagementResourceAssociationsCmd).Standalone()

		partnercentralSelling_listEngagementResourceAssociationsCmd.Flags().String("catalog", "", "Specifies the catalog in which to search for engagement-resource associations.")
		partnercentralSelling_listEngagementResourceAssociationsCmd.Flags().String("created-by", "", "Filters the response to include only snapshots of resources owned by the specified AWS account ID.")
		partnercentralSelling_listEngagementResourceAssociationsCmd.Flags().String("engagement-identifier", "", "Filters the results to include only associations related to the specified engagement.")
		partnercentralSelling_listEngagementResourceAssociationsCmd.Flags().String("max-results", "", "Limits the number of results returned in a single call.")
		partnercentralSelling_listEngagementResourceAssociationsCmd.Flags().String("next-token", "", "A token used for pagination of results.")
		partnercentralSelling_listEngagementResourceAssociationsCmd.Flags().String("resource-identifier", "", "Filters the results to include only associations with the specified resource.")
		partnercentralSelling_listEngagementResourceAssociationsCmd.Flags().String("resource-type", "", "Filters the results to include only associations with resources of the specified type.")
		partnercentralSelling_listEngagementResourceAssociationsCmd.MarkFlagRequired("catalog")
	})
	partnercentralSellingCmd.AddCommand(partnercentralSelling_listEngagementResourceAssociationsCmd)
}
