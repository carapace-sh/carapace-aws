package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteRoleAliasCmd = &cobra.Command{
	Use:   "delete-role-alias",
	Short: "Deletes a role alias\n\nRequires permission to access the [DeleteRoleAlias](https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions) action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteRoleAliasCmd).Standalone()

	iot_deleteRoleAliasCmd.Flags().String("role-alias", "", "The role alias to delete.")
	iot_deleteRoleAliasCmd.MarkFlagRequired("role-alias")
	iotCmd.AddCommand(iot_deleteRoleAliasCmd)
}
