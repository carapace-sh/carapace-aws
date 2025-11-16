package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_createSnapshotCopyConfigurationCmd = &cobra.Command{
	Use:   "create-snapshot-copy-configuration",
	Short: "Creates a snapshot copy configuration that lets you copy snapshots to another Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_createSnapshotCopyConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_createSnapshotCopyConfigurationCmd).Standalone()

		redshiftServerless_createSnapshotCopyConfigurationCmd.Flags().String("destination-kms-key-id", "", "The KMS key to use to encrypt your snapshots in the destination Amazon Web Services Region.")
		redshiftServerless_createSnapshotCopyConfigurationCmd.Flags().String("destination-region", "", "The destination Amazon Web Services Region that you want to copy snapshots to.")
		redshiftServerless_createSnapshotCopyConfigurationCmd.Flags().String("namespace-name", "", "The name of the namespace to copy snapshots from.")
		redshiftServerless_createSnapshotCopyConfigurationCmd.Flags().String("snapshot-retention-period", "", "The retention period of the snapshots that you copy to the destination Amazon Web Services Region.")
		redshiftServerless_createSnapshotCopyConfigurationCmd.MarkFlagRequired("destination-region")
		redshiftServerless_createSnapshotCopyConfigurationCmd.MarkFlagRequired("namespace-name")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_createSnapshotCopyConfigurationCmd)
}
