package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_acceptEngagementInvitationCmd = &cobra.Command{
	Use:   "accept-engagement-invitation",
	Short: "Use the `AcceptEngagementInvitation` action to accept an engagement invitation shared by AWS.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_acceptEngagementInvitationCmd).Standalone()

	partnercentralSelling_acceptEngagementInvitationCmd.Flags().String("catalog", "", "The `CatalogType` parameter specifies the catalog associated with the engagement invitation.")
	partnercentralSelling_acceptEngagementInvitationCmd.Flags().String("identifier", "", "The `Identifier` parameter in the `AcceptEngagementInvitationRequest` specifies the unique identifier of the `EngagementInvitation` to be accepted.")
	partnercentralSelling_acceptEngagementInvitationCmd.MarkFlagRequired("catalog")
	partnercentralSelling_acceptEngagementInvitationCmd.MarkFlagRequired("identifier")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_acceptEngagementInvitationCmd)
}
