package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_getEngagementInvitationCmd = &cobra.Command{
	Use:   "get-engagement-invitation",
	Short: "Retrieves the details of an engagement invitation shared by AWS with a partner.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_getEngagementInvitationCmd).Standalone()

	partnercentralSelling_getEngagementInvitationCmd.Flags().String("catalog", "", "Specifies the catalog associated with the request.")
	partnercentralSelling_getEngagementInvitationCmd.Flags().String("identifier", "", "Specifies the unique identifier for the retrieved engagement invitation.")
	partnercentralSelling_getEngagementInvitationCmd.MarkFlagRequired("catalog")
	partnercentralSelling_getEngagementInvitationCmd.MarkFlagRequired("identifier")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_getEngagementInvitationCmd)
}
