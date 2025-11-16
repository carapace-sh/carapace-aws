package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_createSyncJobCmd = &cobra.Command{
	Use:   "create-sync-job",
	Short: "This action creates a SyncJob.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_createSyncJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_createSyncJobCmd).Standalone()

		iottwinmaker_createSyncJobCmd.Flags().String("sync-role", "", "The SyncJob IAM role.")
		iottwinmaker_createSyncJobCmd.Flags().String("sync-source", "", "The sync source.")
		iottwinmaker_createSyncJobCmd.Flags().String("tags", "", "The SyncJob tags.")
		iottwinmaker_createSyncJobCmd.Flags().String("workspace-id", "", "The workspace ID.")
		iottwinmaker_createSyncJobCmd.MarkFlagRequired("sync-role")
		iottwinmaker_createSyncJobCmd.MarkFlagRequired("sync-source")
		iottwinmaker_createSyncJobCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_createSyncJobCmd)
}
