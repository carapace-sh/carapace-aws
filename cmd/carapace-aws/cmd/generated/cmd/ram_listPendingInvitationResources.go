package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_listPendingInvitationResourcesCmd = &cobra.Command{
	Use:   "list-pending-invitation-resources",
	Short: "Lists the resources in a resource share that is shared with you but for which the invitation is still `PENDING`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_listPendingInvitationResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_listPendingInvitationResourcesCmd).Standalone()

		ram_listPendingInvitationResourcesCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included on each page of the response.")
		ram_listPendingInvitationResourcesCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		ram_listPendingInvitationResourcesCmd.Flags().String("resource-region-scope", "", "Specifies that you want the results to include only resources that have the specified scope.")
		ram_listPendingInvitationResourcesCmd.Flags().String("resource-share-invitation-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the invitation.")
		ram_listPendingInvitationResourcesCmd.MarkFlagRequired("resource-share-invitation-arn")
	})
	ramCmd.AddCommand(ram_listPendingInvitationResourcesCmd)
}
