package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_deletePermissionCmd = &cobra.Command{
	Use:   "delete-permission",
	Short: "Deletes the specified customer managed permission in the Amazon Web Services Region in which you call this operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_deletePermissionCmd).Standalone()

	ram_deletePermissionCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ram_deletePermissionCmd.Flags().String("permission-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the customer managed permission that you want to delete.")
	ram_deletePermissionCmd.MarkFlagRequired("permission-arn")
	ramCmd.AddCommand(ram_deletePermissionCmd)
}
