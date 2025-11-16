package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_createPermissionCmd = &cobra.Command{
	Use:   "create-permission",
	Short: "Creates a customer managed permission for a specified resource type that you can attach to resource shares.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_createPermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_createPermissionCmd).Standalone()

		ram_createPermissionCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ram_createPermissionCmd.Flags().String("name", "", "Specifies the name of the customer managed permission.")
		ram_createPermissionCmd.Flags().String("policy-template", "", "A string in JSON format string that contains the following elements of a resource-based policy:")
		ram_createPermissionCmd.Flags().String("resource-type", "", "Specifies the name of the resource type that this customer managed permission applies to.")
		ram_createPermissionCmd.Flags().String("tags", "", "Specifies a list of one or more tag key and value pairs to attach to the permission.")
		ram_createPermissionCmd.MarkFlagRequired("name")
		ram_createPermissionCmd.MarkFlagRequired("policy-template")
		ram_createPermissionCmd.MarkFlagRequired("resource-type")
	})
	ramCmd.AddCommand(ram_createPermissionCmd)
}
