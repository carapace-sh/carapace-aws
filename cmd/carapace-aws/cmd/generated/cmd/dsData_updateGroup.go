package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsData_updateGroupCmd = &cobra.Command{
	Use:   "update-group",
	Short: "Updates group information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsData_updateGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsData_updateGroupCmd).Standalone()

		dsData_updateGroupCmd.Flags().String("client-token", "", "A unique and case-sensitive identifier that you provide to make sure the idempotency of the request, so multiple identical calls have the same effect as one single call.")
		dsData_updateGroupCmd.Flags().String("directory-id", "", "The identifier (ID) of the directory that's associated with the group.")
		dsData_updateGroupCmd.Flags().String("group-scope", "", "The scope of the AD group.")
		dsData_updateGroupCmd.Flags().String("group-type", "", "The AD group type.")
		dsData_updateGroupCmd.Flags().String("other-attributes", "", "An expression that defines one or more attributes with the data type and the value of each attribute.")
		dsData_updateGroupCmd.Flags().String("samaccount-name", "", "The name of the group.")
		dsData_updateGroupCmd.Flags().String("update-type", "", "The type of update to be performed.")
		dsData_updateGroupCmd.MarkFlagRequired("directory-id")
		dsData_updateGroupCmd.MarkFlagRequired("samaccount-name")
	})
	dsDataCmd.AddCommand(dsData_updateGroupCmd)
}
