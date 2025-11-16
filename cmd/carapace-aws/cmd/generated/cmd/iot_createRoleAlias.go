package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createRoleAliasCmd = &cobra.Command{
	Use:   "create-role-alias",
	Short: "Creates a role alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createRoleAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_createRoleAliasCmd).Standalone()

		iot_createRoleAliasCmd.Flags().String("credential-duration-seconds", "", "How long (in seconds) the credentials will be valid.")
		iot_createRoleAliasCmd.Flags().String("role-alias", "", "The role alias that points to a role ARN.")
		iot_createRoleAliasCmd.Flags().String("role-arn", "", "The role ARN.")
		iot_createRoleAliasCmd.Flags().String("tags", "", "Metadata which can be used to manage the role alias.")
		iot_createRoleAliasCmd.MarkFlagRequired("role-alias")
		iot_createRoleAliasCmd.MarkFlagRequired("role-arn")
	})
	iotCmd.AddCommand(iot_createRoleAliasCmd)
}
