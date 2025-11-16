package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsData_createGroupCmd = &cobra.Command{
	Use:   "create-group",
	Short: "Creates a new group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsData_createGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsData_createGroupCmd).Standalone()

		dsData_createGroupCmd.Flags().String("client-token", "", "A unique and case-sensitive identifier that you provide to make sure the idempotency of the request, so multiple identical calls have the same effect as one single call.")
		dsData_createGroupCmd.Flags().String("directory-id", "", "The identifier (ID) of the directory that's associated with the group.")
		dsData_createGroupCmd.Flags().String("group-scope", "", "The scope of the AD group.")
		dsData_createGroupCmd.Flags().String("group-type", "", "The AD group type.")
		dsData_createGroupCmd.Flags().String("other-attributes", "", "An expression that defines one or more attributes with the data type and value of each attribute.")
		dsData_createGroupCmd.Flags().String("samaccount-name", "", "The name of the group.")
		dsData_createGroupCmd.MarkFlagRequired("directory-id")
		dsData_createGroupCmd.MarkFlagRequired("samaccount-name")
	})
	dsDataCmd.AddCommand(dsData_createGroupCmd)
}
