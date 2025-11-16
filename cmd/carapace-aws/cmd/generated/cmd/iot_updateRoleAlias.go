package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateRoleAliasCmd = &cobra.Command{
	Use:   "update-role-alias",
	Short: "Updates a role alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateRoleAliasCmd).Standalone()

	iot_updateRoleAliasCmd.Flags().String("credential-duration-seconds", "", "The number of seconds the credential will be valid.")
	iot_updateRoleAliasCmd.Flags().String("role-alias", "", "The role alias to update.")
	iot_updateRoleAliasCmd.Flags().String("role-arn", "", "The role ARN.")
	iot_updateRoleAliasCmd.MarkFlagRequired("role-alias")
	iotCmd.AddCommand(iot_updateRoleAliasCmd)
}
