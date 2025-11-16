package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_getSyncJobCmd = &cobra.Command{
	Use:   "get-sync-job",
	Short: "Gets the SyncJob.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_getSyncJobCmd).Standalone()

	iottwinmaker_getSyncJobCmd.Flags().String("sync-source", "", "The sync source.")
	iottwinmaker_getSyncJobCmd.Flags().String("workspace-id", "", "The workspace ID.")
	iottwinmaker_getSyncJobCmd.MarkFlagRequired("sync-source")
	iottwinmakerCmd.AddCommand(iottwinmaker_getSyncJobCmd)
}
