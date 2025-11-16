package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_getResourceShareInvitationsCmd = &cobra.Command{
	Use:   "get-resource-share-invitations",
	Short: "Retrieves details about invitations that you have received for resource shares.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_getResourceShareInvitationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_getResourceShareInvitationsCmd).Standalone()

		ram_getResourceShareInvitationsCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included on each page of the response.")
		ram_getResourceShareInvitationsCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		ram_getResourceShareInvitationsCmd.Flags().String("resource-share-arns", "", "Specifies that you want details about invitations only for the resource shares described by this list of [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)")
		ram_getResourceShareInvitationsCmd.Flags().String("resource-share-invitation-arns", "", "Specifies the [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share invitations you want information about.")
	})
	ramCmd.AddCommand(ram_getResourceShareInvitationsCmd)
}
