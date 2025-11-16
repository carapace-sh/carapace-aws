package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_replacePermissionAssociationsCmd = &cobra.Command{
	Use:   "replace-permission-associations",
	Short: "Updates all resource shares that use a managed permission to a different managed permission.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_replacePermissionAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_replacePermissionAssociationsCmd).Standalone()

		ram_replacePermissionAssociationsCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ram_replacePermissionAssociationsCmd.Flags().String("from-permission-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the managed permission that you want to replace.")
		ram_replacePermissionAssociationsCmd.Flags().String("from-permission-version", "", "Specifies that you want to updated the permissions for only those resource shares that use the specified version of the managed permission.")
		ram_replacePermissionAssociationsCmd.Flags().String("to-permission-arn", "", "Specifies the ARN of the managed permission that you want to associate with resource shares in place of the one specified by `fromPerssionArn` and `fromPermissionVersion`.")
		ram_replacePermissionAssociationsCmd.MarkFlagRequired("from-permission-arn")
		ram_replacePermissionAssociationsCmd.MarkFlagRequired("to-permission-arn")
	})
	ramCmd.AddCommand(ram_replacePermissionAssociationsCmd)
}
