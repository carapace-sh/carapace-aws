package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_getRepositorySyncStatusCmd = &cobra.Command{
	Use:   "get-repository-sync-status",
	Short: "Get the sync status of a repository used for Proton template sync.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_getRepositorySyncStatusCmd).Standalone()

	proton_getRepositorySyncStatusCmd.Flags().String("branch", "", "The repository branch.")
	proton_getRepositorySyncStatusCmd.Flags().String("repository-name", "", "The repository name.")
	proton_getRepositorySyncStatusCmd.Flags().String("repository-provider", "", "The repository provider.")
	proton_getRepositorySyncStatusCmd.Flags().String("sync-type", "", "The repository sync type.")
	proton_getRepositorySyncStatusCmd.MarkFlagRequired("branch")
	proton_getRepositorySyncStatusCmd.MarkFlagRequired("repository-name")
	proton_getRepositorySyncStatusCmd.MarkFlagRequired("repository-provider")
	proton_getRepositorySyncStatusCmd.MarkFlagRequired("sync-type")
	protonCmd.AddCommand(proton_getRepositorySyncStatusCmd)
}
