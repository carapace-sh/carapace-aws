package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_updateSnapshotScheduleCmd = &cobra.Command{
	Use:   "update-snapshot-schedule",
	Short: "Updates a snapshot schedule configured for a gateway volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_updateSnapshotScheduleCmd).Standalone()

	storagegateway_updateSnapshotScheduleCmd.Flags().String("description", "", "Optional description of the snapshot that overwrites the existing description.")
	storagegateway_updateSnapshotScheduleCmd.Flags().String("recurrence-in-hours", "", "Frequency of snapshots.")
	storagegateway_updateSnapshotScheduleCmd.Flags().String("start-at", "", "The hour of the day at which the snapshot schedule begins represented as *hh*, where *hh* is the hour (0 to 23).")
	storagegateway_updateSnapshotScheduleCmd.Flags().String("tags", "", "A list of up to 50 tags that can be assigned to a snapshot.")
	storagegateway_updateSnapshotScheduleCmd.Flags().String("volume-arn", "", "The Amazon Resource Name (ARN) of the volume.")
	storagegateway_updateSnapshotScheduleCmd.MarkFlagRequired("recurrence-in-hours")
	storagegateway_updateSnapshotScheduleCmd.MarkFlagRequired("start-at")
	storagegateway_updateSnapshotScheduleCmd.MarkFlagRequired("volume-arn")
	storagegatewayCmd.AddCommand(storagegateway_updateSnapshotScheduleCmd)
}
