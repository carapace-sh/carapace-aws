package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_deleteSnapshotCopyConfigurationCmd = &cobra.Command{
	Use:   "delete-snapshot-copy-configuration",
	Short: "Deletes a snapshot copy configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_deleteSnapshotCopyConfigurationCmd).Standalone()

	redshiftServerless_deleteSnapshotCopyConfigurationCmd.Flags().String("snapshot-copy-configuration-id", "", "The ID of the snapshot copy configuration to delete.")
	redshiftServerless_deleteSnapshotCopyConfigurationCmd.MarkFlagRequired("snapshot-copy-configuration-id")
	redshiftServerlessCmd.AddCommand(redshiftServerless_deleteSnapshotCopyConfigurationCmd)
}
