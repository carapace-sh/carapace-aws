package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_createEngagementInvitationCmd = &cobra.Command{
	Use:   "create-engagement-invitation",
	Short: "This action creates an invitation from a sender to a single receiver to join an engagement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_createEngagementInvitationCmd).Standalone()

	partnercentralSelling_createEngagementInvitationCmd.Flags().String("catalog", "", "Specifies the catalog related to the engagement.")
	partnercentralSelling_createEngagementInvitationCmd.Flags().String("client-token", "", "Specifies a unique, client-generated UUID to ensure that the request is handled exactly once.")
	partnercentralSelling_createEngagementInvitationCmd.Flags().String("engagement-identifier", "", "The unique identifier of the `Engagement` associated with the invitation.")
	partnercentralSelling_createEngagementInvitationCmd.Flags().String("invitation", "", "The `Invitation` object all information necessary to initiate an engagement invitation to a partner.")
	partnercentralSelling_createEngagementInvitationCmd.MarkFlagRequired("catalog")
	partnercentralSelling_createEngagementInvitationCmd.MarkFlagRequired("client-token")
	partnercentralSelling_createEngagementInvitationCmd.MarkFlagRequired("engagement-identifier")
	partnercentralSelling_createEngagementInvitationCmd.MarkFlagRequired("invitation")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_createEngagementInvitationCmd)
}
