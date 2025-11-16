package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeSnapshotScheduleCmd = &cobra.Command{
	Use:   "describe-snapshot-schedule",
	Short: "Describes the snapshot schedule for the specified gateway volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeSnapshotScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_describeSnapshotScheduleCmd).Standalone()

		storagegateway_describeSnapshotScheduleCmd.Flags().String("volume-arn", "", "The Amazon Resource Name (ARN) of the volume.")
		storagegateway_describeSnapshotScheduleCmd.MarkFlagRequired("volume-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_describeSnapshotScheduleCmd)
}
