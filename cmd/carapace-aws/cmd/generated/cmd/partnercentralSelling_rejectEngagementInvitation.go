package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_rejectEngagementInvitationCmd = &cobra.Command{
	Use:   "reject-engagement-invitation",
	Short: "This action rejects an `EngagementInvitation` that AWS shared.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_rejectEngagementInvitationCmd).Standalone()

	partnercentralSelling_rejectEngagementInvitationCmd.Flags().String("catalog", "", "This is the catalog that's associated with the engagement invitation.")
	partnercentralSelling_rejectEngagementInvitationCmd.Flags().String("identifier", "", "This is the unique identifier of the rejected `EngagementInvitation`.")
	partnercentralSelling_rejectEngagementInvitationCmd.Flags().String("rejection-reason", "", "This describes the reason for rejecting the engagement invitation, which helps AWS track usage patterns.")
	partnercentralSelling_rejectEngagementInvitationCmd.MarkFlagRequired("catalog")
	partnercentralSelling_rejectEngagementInvitationCmd.MarkFlagRequired("identifier")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_rejectEngagementInvitationCmd)
}
