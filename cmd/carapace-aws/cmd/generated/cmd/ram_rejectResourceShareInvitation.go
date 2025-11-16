package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_rejectResourceShareInvitationCmd = &cobra.Command{
	Use:   "reject-resource-share-invitation",
	Short: "Rejects an invitation to a resource share from another Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_rejectResourceShareInvitationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_rejectResourceShareInvitationCmd).Standalone()

		ram_rejectResourceShareInvitationCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ram_rejectResourceShareInvitationCmd.Flags().String("resource-share-invitation-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the invitation that you want to reject.")
		ram_rejectResourceShareInvitationCmd.MarkFlagRequired("resource-share-invitation-arn")
	})
	ramCmd.AddCommand(ram_rejectResourceShareInvitationCmd)
}
