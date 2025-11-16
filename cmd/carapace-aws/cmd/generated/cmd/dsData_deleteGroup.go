package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsData_deleteGroupCmd = &cobra.Command{
	Use:   "delete-group",
	Short: "Deletes a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsData_deleteGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsData_deleteGroupCmd).Standalone()

		dsData_deleteGroupCmd.Flags().String("client-token", "", "A unique and case-sensitive identifier that you provide to make sure the idempotency of the request, so multiple identical calls have the same effect as one single call.")
		dsData_deleteGroupCmd.Flags().String("directory-id", "", "The identifier (ID) of the directory that's associated with the group.")
		dsData_deleteGroupCmd.Flags().String("samaccount-name", "", "The name of the group.")
		dsData_deleteGroupCmd.MarkFlagRequired("directory-id")
		dsData_deleteGroupCmd.MarkFlagRequired("samaccount-name")
	})
	dsDataCmd.AddCommand(dsData_deleteGroupCmd)
}
