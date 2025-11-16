package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listProtectedResourcesCmd = &cobra.Command{
	Use:   "list-protected-resources",
	Short: "Returns an array of resources successfully backed up by Backup, including the time the resource was saved, an Amazon Resource Name (ARN) of the resource, and a resource type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listProtectedResourcesCmd).Standalone()

	backup_listProtectedResourcesCmd.Flags().String("max-results", "", "The maximum number of items to be returned.")
	backup_listProtectedResourcesCmd.Flags().String("next-token", "", "The next item following a partial list of returned items.")
	backupCmd.AddCommand(backup_listProtectedResourcesCmd)
}
