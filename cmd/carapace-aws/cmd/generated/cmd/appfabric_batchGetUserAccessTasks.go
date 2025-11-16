package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_batchGetUserAccessTasksCmd = &cobra.Command{
	Use:   "batch-get-user-access-tasks",
	Short: "Gets user access details in a batch request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_batchGetUserAccessTasksCmd).Standalone()

	appfabric_batchGetUserAccessTasksCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.")
	appfabric_batchGetUserAccessTasksCmd.Flags().String("task-id-list", "", "The tasks IDs to use for the request.")
	appfabric_batchGetUserAccessTasksCmd.MarkFlagRequired("app-bundle-identifier")
	appfabric_batchGetUserAccessTasksCmd.MarkFlagRequired("task-id-list")
	appfabricCmd.AddCommand(appfabric_batchGetUserAccessTasksCmd)
}
