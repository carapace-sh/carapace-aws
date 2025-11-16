package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_deletePermissionVersionCmd = &cobra.Command{
	Use:   "delete-permission-version",
	Short: "Deletes one version of a customer managed permission.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_deletePermissionVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_deletePermissionVersionCmd).Standalone()

		ram_deletePermissionVersionCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ram_deletePermissionVersionCmd.Flags().String("permission-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the permission with the version you want to delete.")
		ram_deletePermissionVersionCmd.Flags().String("permission-version", "", "Specifies the version number to delete.")
		ram_deletePermissionVersionCmd.MarkFlagRequired("permission-arn")
		ram_deletePermissionVersionCmd.MarkFlagRequired("permission-version")
	})
	ramCmd.AddCommand(ram_deletePermissionVersionCmd)
}
