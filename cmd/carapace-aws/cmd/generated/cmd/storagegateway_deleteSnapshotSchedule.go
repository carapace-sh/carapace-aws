package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_deleteSnapshotScheduleCmd = &cobra.Command{
	Use:   "delete-snapshot-schedule",
	Short: "Deletes a snapshot of a volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_deleteSnapshotScheduleCmd).Standalone()

	storagegateway_deleteSnapshotScheduleCmd.Flags().String("volume-arn", "", "The volume which snapshot schedule to delete.")
	storagegateway_deleteSnapshotScheduleCmd.MarkFlagRequired("volume-arn")
	storagegatewayCmd.AddCommand(storagegateway_deleteSnapshotScheduleCmd)
}
