package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_getPermissionCmd = &cobra.Command{
	Use:   "get-permission",
	Short: "Retrieves the contents of a managed permission in JSON format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_getPermissionCmd).Standalone()

	ram_getPermissionCmd.Flags().String("permission-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the permission whose contents you want to retrieve.")
	ram_getPermissionCmd.Flags().String("permission-version", "", "Specifies the version number of the RAM permission to retrieve.")
	ram_getPermissionCmd.MarkFlagRequired("permission-arn")
	ramCmd.AddCommand(ram_getPermissionCmd)
}
