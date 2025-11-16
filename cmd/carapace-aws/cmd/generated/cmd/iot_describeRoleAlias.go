package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeRoleAliasCmd = &cobra.Command{
	Use:   "describe-role-alias",
	Short: "Describes a role alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeRoleAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeRoleAliasCmd).Standalone()

		iot_describeRoleAliasCmd.Flags().String("role-alias", "", "The role alias to describe.")
		iot_describeRoleAliasCmd.MarkFlagRequired("role-alias")
	})
	iotCmd.AddCommand(iot_describeRoleAliasCmd)
}
