package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_listEngagementInvitationsCmd = &cobra.Command{
	Use:   "list-engagement-invitations",
	Short: "Retrieves a list of engagement invitations sent to the partner.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_listEngagementInvitationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(partnercentralSelling_listEngagementInvitationsCmd).Standalone()

		partnercentralSelling_listEngagementInvitationsCmd.Flags().String("catalog", "", "Specifies the catalog from which to list the engagement invitations.")
		partnercentralSelling_listEngagementInvitationsCmd.Flags().String("engagement-identifier", "", "Retrieves a list of engagement invitation summaries based on specified filters.")
		partnercentralSelling_listEngagementInvitationsCmd.Flags().String("max-results", "", "Specifies the maximum number of engagement invitations to return in the response.")
		partnercentralSelling_listEngagementInvitationsCmd.Flags().String("next-token", "", "A pagination token used to retrieve additional pages of results when the response to a previous request was truncated.")
		partnercentralSelling_listEngagementInvitationsCmd.Flags().String("participant-type", "", "Specifies the type of participant for which to list engagement invitations.")
		partnercentralSelling_listEngagementInvitationsCmd.Flags().String("payload-type", "", "Defines the type of payload associated with the engagement invitations to be listed.")
		partnercentralSelling_listEngagementInvitationsCmd.Flags().String("sender-aws-account-id", "", "List of sender AWS account IDs to filter the invitations.")
		partnercentralSelling_listEngagementInvitationsCmd.Flags().String("sort", "", "Specifies the sorting options for listing engagement invitations.")
		partnercentralSelling_listEngagementInvitationsCmd.Flags().String("status", "", "Status values to filter the invitations.")
		partnercentralSelling_listEngagementInvitationsCmd.MarkFlagRequired("catalog")
		partnercentralSelling_listEngagementInvitationsCmd.MarkFlagRequired("participant-type")
	})
	partnercentralSellingCmd.AddCommand(partnercentralSelling_listEngagementInvitationsCmd)
}
