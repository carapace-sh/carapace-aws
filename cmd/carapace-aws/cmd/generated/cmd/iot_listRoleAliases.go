package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listRoleAliasesCmd = &cobra.Command{
	Use:   "list-role-aliases",
	Short: "Lists the role aliases registered in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listRoleAliasesCmd).Standalone()

	iot_listRoleAliasesCmd.Flags().String("ascending-order", "", "Return the list of role aliases in ascending alphabetical order.")
	iot_listRoleAliasesCmd.Flags().String("marker", "", "A marker used to get the next set of results.")
	iot_listRoleAliasesCmd.Flags().String("page-size", "", "The maximum number of results to return at one time.")
	iotCmd.AddCommand(iot_listRoleAliasesCmd)
}
