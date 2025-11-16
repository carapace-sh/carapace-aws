package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_startUserAccessTasksCmd = &cobra.Command{
	Use:   "start-user-access-tasks",
	Short: "Starts the tasks to search user access status for a specific email address.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_startUserAccessTasksCmd).Standalone()

	appfabric_startUserAccessTasksCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.")
	appfabric_startUserAccessTasksCmd.Flags().String("email", "", "The email address of the target user.")
	appfabric_startUserAccessTasksCmd.MarkFlagRequired("app-bundle-identifier")
	appfabric_startUserAccessTasksCmd.MarkFlagRequired("email")
	appfabricCmd.AddCommand(appfabric_startUserAccessTasksCmd)
}
