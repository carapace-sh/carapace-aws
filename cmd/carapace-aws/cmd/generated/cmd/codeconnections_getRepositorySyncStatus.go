package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_getRepositorySyncStatusCmd = &cobra.Command{
	Use:   "get-repository-sync-status",
	Short: "Returns details about the sync status for a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_getRepositorySyncStatusCmd).Standalone()

	codeconnections_getRepositorySyncStatusCmd.Flags().String("branch", "", "The branch of the repository link for the requested repository sync status.")
	codeconnections_getRepositorySyncStatusCmd.Flags().String("repository-link-id", "", "The repository link ID for the requested repository sync status.")
	codeconnections_getRepositorySyncStatusCmd.Flags().String("sync-type", "", "The sync type of the requested sync status.")
	codeconnections_getRepositorySyncStatusCmd.MarkFlagRequired("branch")
	codeconnections_getRepositorySyncStatusCmd.MarkFlagRequired("repository-link-id")
	codeconnections_getRepositorySyncStatusCmd.MarkFlagRequired("sync-type")
	codeconnectionsCmd.AddCommand(codeconnections_getRepositorySyncStatusCmd)
}
