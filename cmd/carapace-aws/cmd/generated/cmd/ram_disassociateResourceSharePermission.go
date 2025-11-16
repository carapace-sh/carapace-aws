package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_disassociateResourceSharePermissionCmd = &cobra.Command{
	Use:   "disassociate-resource-share-permission",
	Short: "Removes a managed permission from a resource share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_disassociateResourceSharePermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_disassociateResourceSharePermissionCmd).Standalone()

		ram_disassociateResourceSharePermissionCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ram_disassociateResourceSharePermissionCmd.Flags().String("permission-arn", "", "The [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the managed permission to disassociate from the resource share.")
		ram_disassociateResourceSharePermissionCmd.Flags().String("resource-share-arn", "", "The [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share that you want to remove the managed permission from.")
		ram_disassociateResourceSharePermissionCmd.MarkFlagRequired("permission-arn")
		ram_disassociateResourceSharePermissionCmd.MarkFlagRequired("resource-share-arn")
	})
	ramCmd.AddCommand(ram_disassociateResourceSharePermissionCmd)
}
