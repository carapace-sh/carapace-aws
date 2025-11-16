package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_setDefaultPermissionVersionCmd = &cobra.Command{
	Use:   "set-default-permission-version",
	Short: "Designates the specified version number as the default version for the specified customer managed permission.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_setDefaultPermissionVersionCmd).Standalone()

	ram_setDefaultPermissionVersionCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ram_setDefaultPermissionVersionCmd.Flags().String("permission-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the customer managed permission whose default version you want to change.")
	ram_setDefaultPermissionVersionCmd.Flags().String("permission-version", "", "Specifies the version number that you want to designate as the default for customer managed permission.")
	ram_setDefaultPermissionVersionCmd.MarkFlagRequired("permission-arn")
	ram_setDefaultPermissionVersionCmd.MarkFlagRequired("permission-version")
	ramCmd.AddCommand(ram_setDefaultPermissionVersionCmd)
}
