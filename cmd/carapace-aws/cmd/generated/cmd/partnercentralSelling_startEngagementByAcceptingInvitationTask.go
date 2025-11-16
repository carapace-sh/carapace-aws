package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_startEngagementByAcceptingInvitationTaskCmd = &cobra.Command{
	Use:   "start-engagement-by-accepting-invitation-task",
	Short: "This action starts the engagement by accepting an `EngagementInvitation`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_startEngagementByAcceptingInvitationTaskCmd).Standalone()

	partnercentralSelling_startEngagementByAcceptingInvitationTaskCmd.Flags().String("catalog", "", "Specifies the catalog related to the task.")
	partnercentralSelling_startEngagementByAcceptingInvitationTaskCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier provided by the client that helps to ensure the idempotency of the request.")
	partnercentralSelling_startEngagementByAcceptingInvitationTaskCmd.Flags().String("identifier", "", "Specifies the unique identifier of the `EngagementInvitation` to be accepted.")
	partnercentralSelling_startEngagementByAcceptingInvitationTaskCmd.Flags().String("tags", "", "A map of the key-value pairs of the tag or tags to assign.")
	partnercentralSelling_startEngagementByAcceptingInvitationTaskCmd.MarkFlagRequired("catalog")
	partnercentralSelling_startEngagementByAcceptingInvitationTaskCmd.MarkFlagRequired("client-token")
	partnercentralSelling_startEngagementByAcceptingInvitationTaskCmd.MarkFlagRequired("identifier")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_startEngagementByAcceptingInvitationTaskCmd)
}
