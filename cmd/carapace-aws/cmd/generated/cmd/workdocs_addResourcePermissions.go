package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_addResourcePermissionsCmd = &cobra.Command{
	Use:   "add-resource-permissions",
	Short: "Creates a set of permissions for the specified folder or document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_addResourcePermissionsCmd).Standalone()

	workdocs_addResourcePermissionsCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_addResourcePermissionsCmd.Flags().String("notification-options", "", "The notification options.")
	workdocs_addResourcePermissionsCmd.Flags().String("principals", "", "The users, groups, or organization being granted permission.")
	workdocs_addResourcePermissionsCmd.Flags().String("resource-id", "", "The ID of the resource.")
	workdocs_addResourcePermissionsCmd.MarkFlagRequired("principals")
	workdocs_addResourcePermissionsCmd.MarkFlagRequired("resource-id")
	workdocsCmd.AddCommand(workdocs_addResourcePermissionsCmd)
}
