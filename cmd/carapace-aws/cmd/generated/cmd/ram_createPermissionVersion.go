package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_createPermissionVersionCmd = &cobra.Command{
	Use:   "create-permission-version",
	Short: "Creates a new version of the specified customer managed permission.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_createPermissionVersionCmd).Standalone()

	ram_createPermissionVersionCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ram_createPermissionVersionCmd.Flags().String("permission-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the customer managed permission you're creating a new version for.")
	ram_createPermissionVersionCmd.Flags().String("policy-template", "", "A string in JSON format string that contains the following elements of a resource-based policy:")
	ram_createPermissionVersionCmd.MarkFlagRequired("permission-arn")
	ram_createPermissionVersionCmd.MarkFlagRequired("policy-template")
	ramCmd.AddCommand(ram_createPermissionVersionCmd)
}
