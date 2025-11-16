package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_deleteSyncJobCmd = &cobra.Command{
	Use:   "delete-sync-job",
	Short: "Delete the SyncJob.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_deleteSyncJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_deleteSyncJobCmd).Standalone()

		iottwinmaker_deleteSyncJobCmd.Flags().String("sync-source", "", "The sync source.")
		iottwinmaker_deleteSyncJobCmd.Flags().String("workspace-id", "", "The workspace ID.")
		iottwinmaker_deleteSyncJobCmd.MarkFlagRequired("sync-source")
		iottwinmaker_deleteSyncJobCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_deleteSyncJobCmd)
}
