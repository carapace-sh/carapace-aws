package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deleteSnapshotScheduleCmd = &cobra.Command{
	Use:   "delete-snapshot-schedule",
	Short: "Deletes a snapshot schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deleteSnapshotScheduleCmd).Standalone()

	redshift_deleteSnapshotScheduleCmd.Flags().String("schedule-identifier", "", "A unique identifier of the snapshot schedule to delete.")
	redshift_deleteSnapshotScheduleCmd.MarkFlagRequired("schedule-identifier")
	redshiftCmd.AddCommand(redshift_deleteSnapshotScheduleCmd)
}
