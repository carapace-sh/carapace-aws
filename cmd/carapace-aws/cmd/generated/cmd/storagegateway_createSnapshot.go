package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_createSnapshotCmd = &cobra.Command{
	Use:   "create-snapshot",
	Short: "Initiates a snapshot of a volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_createSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_createSnapshotCmd).Standalone()

		storagegateway_createSnapshotCmd.Flags().String("snapshot-description", "", "Textual description of the snapshot that appears in the Amazon EC2 console, Elastic Block Store snapshots panel in the **Description** field, and in the Storage Gateway snapshot **Details** pane, **Description** field.")
		storagegateway_createSnapshotCmd.Flags().String("tags", "", "A list of up to 50 tags that can be assigned to a snapshot.")
		storagegateway_createSnapshotCmd.Flags().String("volume-arn", "", "The Amazon Resource Name (ARN) of the volume.")
		storagegateway_createSnapshotCmd.MarkFlagRequired("snapshot-description")
		storagegateway_createSnapshotCmd.MarkFlagRequired("volume-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_createSnapshotCmd)
}
