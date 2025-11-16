package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_acceptResourceShareInvitationCmd = &cobra.Command{
	Use:   "accept-resource-share-invitation",
	Short: "Accepts an invitation to a resource share from another Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_acceptResourceShareInvitationCmd).Standalone()

	ram_acceptResourceShareInvitationCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ram_acceptResourceShareInvitationCmd.Flags().String("resource-share-invitation-arn", "", "The [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the invitation that you want to accept.")
	ram_acceptResourceShareInvitationCmd.MarkFlagRequired("resource-share-invitation-arn")
	ramCmd.AddCommand(ram_acceptResourceShareInvitationCmd)
}
