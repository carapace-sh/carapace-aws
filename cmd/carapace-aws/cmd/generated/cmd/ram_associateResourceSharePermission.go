package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_associateResourceSharePermissionCmd = &cobra.Command{
	Use:   "associate-resource-share-permission",
	Short: "Adds or replaces the RAM permission for a resource type included in a resource share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_associateResourceSharePermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_associateResourceSharePermissionCmd).Standalone()

		ram_associateResourceSharePermissionCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ram_associateResourceSharePermissionCmd.Flags().Bool("no-replace", false, "Specifies whether the specified permission should replace the existing permission associated with the resource share.")
		ram_associateResourceSharePermissionCmd.Flags().String("permission-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the RAM permission to associate with the resource share.")
		ram_associateResourceSharePermissionCmd.Flags().String("permission-version", "", "Specifies the version of the RAM permission to associate with the resource share.")
		ram_associateResourceSharePermissionCmd.Flags().Bool("replace", false, "Specifies whether the specified permission should replace the existing permission associated with the resource share.")
		ram_associateResourceSharePermissionCmd.Flags().String("resource-share-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share to which you want to add or replace permissions.")
		ram_associateResourceSharePermissionCmd.Flag("no-replace").Hidden = true
		ram_associateResourceSharePermissionCmd.MarkFlagRequired("permission-arn")
		ram_associateResourceSharePermissionCmd.MarkFlagRequired("resource-share-arn")
	})
	ramCmd.AddCommand(ram_associateResourceSharePermissionCmd)
}
