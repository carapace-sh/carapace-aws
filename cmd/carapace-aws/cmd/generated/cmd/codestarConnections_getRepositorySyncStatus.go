package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_getRepositorySyncStatusCmd = &cobra.Command{
	Use:   "get-repository-sync-status",
	Short: "Returns details about the sync status for a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_getRepositorySyncStatusCmd).Standalone()

	codestarConnections_getRepositorySyncStatusCmd.Flags().String("branch", "", "The branch of the repository link for the requested repository sync status.")
	codestarConnections_getRepositorySyncStatusCmd.Flags().String("repository-link-id", "", "The repository link ID for the requested repository sync status.")
	codestarConnections_getRepositorySyncStatusCmd.Flags().String("sync-type", "", "The sync type of the requested sync status.")
	codestarConnections_getRepositorySyncStatusCmd.MarkFlagRequired("branch")
	codestarConnections_getRepositorySyncStatusCmd.MarkFlagRequired("repository-link-id")
	codestarConnections_getRepositorySyncStatusCmd.MarkFlagRequired("sync-type")
	codestarConnectionsCmd.AddCommand(codestarConnections_getRepositorySyncStatusCmd)
}
