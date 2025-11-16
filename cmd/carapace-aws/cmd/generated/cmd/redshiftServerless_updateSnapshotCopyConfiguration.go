package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_updateSnapshotCopyConfigurationCmd = &cobra.Command{
	Use:   "update-snapshot-copy-configuration",
	Short: "Updates a snapshot copy configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_updateSnapshotCopyConfigurationCmd).Standalone()

	redshiftServerless_updateSnapshotCopyConfigurationCmd.Flags().String("snapshot-copy-configuration-id", "", "The ID of the snapshot copy configuration to update.")
	redshiftServerless_updateSnapshotCopyConfigurationCmd.Flags().String("snapshot-retention-period", "", "The new retention period of how long to keep a snapshot in the destination Amazon Web Services Region.")
	redshiftServerless_updateSnapshotCopyConfigurationCmd.MarkFlagRequired("snapshot-copy-configuration-id")
	redshiftServerlessCmd.AddCommand(redshiftServerless_updateSnapshotCopyConfigurationCmd)
}
