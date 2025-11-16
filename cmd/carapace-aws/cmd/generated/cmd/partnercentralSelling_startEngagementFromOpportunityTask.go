package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_startEngagementFromOpportunityTaskCmd = &cobra.Command{
	Use:   "start-engagement-from-opportunity-task",
	Short: "Similar to `StartEngagementByAcceptingInvitationTask`, this action is asynchronous and performs multiple steps before completion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_startEngagementFromOpportunityTaskCmd).Standalone()

	partnercentralSelling_startEngagementFromOpportunityTaskCmd.Flags().String("aws-submission", "", "")
	partnercentralSelling_startEngagementFromOpportunityTaskCmd.Flags().String("catalog", "", "Specifies the catalog in which the engagement is tracked.")
	partnercentralSelling_startEngagementFromOpportunityTaskCmd.Flags().String("client-token", "", "A unique token provided by the client to help ensure the idempotency of the request.")
	partnercentralSelling_startEngagementFromOpportunityTaskCmd.Flags().String("identifier", "", "The unique identifier of the opportunity from which the engagement task is to be initiated.")
	partnercentralSelling_startEngagementFromOpportunityTaskCmd.Flags().String("tags", "", "A map of the key-value pairs of the tag or tags to assign.")
	partnercentralSelling_startEngagementFromOpportunityTaskCmd.MarkFlagRequired("aws-submission")
	partnercentralSelling_startEngagementFromOpportunityTaskCmd.MarkFlagRequired("catalog")
	partnercentralSelling_startEngagementFromOpportunityTaskCmd.MarkFlagRequired("client-token")
	partnercentralSelling_startEngagementFromOpportunityTaskCmd.MarkFlagRequired("identifier")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_startEngagementFromOpportunityTaskCmd)
}
