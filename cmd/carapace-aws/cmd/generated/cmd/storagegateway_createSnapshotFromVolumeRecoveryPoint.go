package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_createSnapshotFromVolumeRecoveryPointCmd = &cobra.Command{
	Use:   "create-snapshot-from-volume-recovery-point",
	Short: "Initiates a snapshot of a gateway from a volume recovery point.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_createSnapshotFromVolumeRecoveryPointCmd).Standalone()

	storagegateway_createSnapshotFromVolumeRecoveryPointCmd.Flags().String("snapshot-description", "", "Textual description of the snapshot that appears in the Amazon EC2 console, Elastic Block Store snapshots panel in the **Description** field, and in the Storage Gateway snapshot **Details** pane, **Description** field.")
	storagegateway_createSnapshotFromVolumeRecoveryPointCmd.Flags().String("tags", "", "A list of up to 50 tags that can be assigned to a snapshot.")
	storagegateway_createSnapshotFromVolumeRecoveryPointCmd.Flags().String("volume-arn", "", "The Amazon Resource Name (ARN) of the iSCSI volume target.")
	storagegateway_createSnapshotFromVolumeRecoveryPointCmd.MarkFlagRequired("snapshot-description")
	storagegateway_createSnapshotFromVolumeRecoveryPointCmd.MarkFlagRequired("volume-arn")
	storagegatewayCmd.AddCommand(storagegateway_createSnapshotFromVolumeRecoveryPointCmd)
}
