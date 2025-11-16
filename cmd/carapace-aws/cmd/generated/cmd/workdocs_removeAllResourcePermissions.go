package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_removeAllResourcePermissionsCmd = &cobra.Command{
	Use:   "remove-all-resource-permissions",
	Short: "Removes all the permissions from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_removeAllResourcePermissionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_removeAllResourcePermissionsCmd).Standalone()

		workdocs_removeAllResourcePermissionsCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
		workdocs_removeAllResourcePermissionsCmd.Flags().String("resource-id", "", "The ID of the resource.")
		workdocs_removeAllResourcePermissionsCmd.MarkFlagRequired("resource-id")
	})
	workdocsCmd.AddCommand(workdocs_removeAllResourcePermissionsCmd)
}
