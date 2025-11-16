package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_removeResourcePermissionCmd = &cobra.Command{
	Use:   "remove-resource-permission",
	Short: "Removes the permission for the specified principal from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_removeResourcePermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_removeResourcePermissionCmd).Standalone()

		workdocs_removeResourcePermissionCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
		workdocs_removeResourcePermissionCmd.Flags().String("principal-id", "", "The principal ID of the resource.")
		workdocs_removeResourcePermissionCmd.Flags().String("principal-type", "", "The principal type of the resource.")
		workdocs_removeResourcePermissionCmd.Flags().String("resource-id", "", "The ID of the resource.")
		workdocs_removeResourcePermissionCmd.MarkFlagRequired("principal-id")
		workdocs_removeResourcePermissionCmd.MarkFlagRequired("resource-id")
	})
	workdocsCmd.AddCommand(workdocs_removeResourcePermissionCmd)
}
