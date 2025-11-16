package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_listPermissionsCmd = &cobra.Command{
	Use:   "list-permissions",
	Short: "Retrieves a list of available RAM permissions that you can use for the supported resource types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_listPermissionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_listPermissionsCmd).Standalone()

		ram_listPermissionsCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included on each page of the response.")
		ram_listPermissionsCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		ram_listPermissionsCmd.Flags().String("permission-type", "", "Specifies that you want to list only permissions of this type:")
		ram_listPermissionsCmd.Flags().String("resource-type", "", "Specifies that you want to list only those permissions that apply to the specified resource type.")
	})
	ramCmd.AddCommand(ram_listPermissionsCmd)
}
